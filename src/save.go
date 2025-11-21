package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/mathrock-xyz/starducc/src/auth"
	"github.com/mathrock-xyz/starducc/src/db"
	"github.com/mathrock-xyz/starducc/src/db/model"
	"github.com/mathrock-xyz/starducc/src/rock"
	"github.com/mathrock-xyz/starducc/src/storage"
	"github.com/redis/go-redis/v9"
)

func save(ctx echo.Context) (err error) {
	id := auth.UserId(ctx)

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return
	}

	descriptor, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer descriptor.Close()

	tx := db.DB.Begin()
	defer tx.Rollback()

	// Cari file berdasarkan (user_id, filename)
	fl := new(model.File)
	err = tx.Where("name = ? AND user_id = ?", fileHeader.Filename, id).First(&fl).Error
	if err != nil {
		return
	}

	// Hash
	hasher := sha256.New()
	if _, err = io.Copy(hasher, descriptor); err != nil {
		return
	}
	descriptor.Seek(0, io.SeekStart)
	hash := fmt.Sprintf("%x", hasher.Sum(nil))

	lastVersion := new(model.FileVersion)
	_ = tx.Where("file_id = ?", fl.ID).Order("version DESC").First(&lastVersion).Error

	if lastVersion.Hash == hash {
		return ctx.JSON(http.StatusOK, map[string]string{"message": "no change"})
	}

	_, err = rock.Rock.Get(context.Background(), hash).Result()
	if err == redis.Nil {
		if _, err = storage.Box.PutObject(context.Background(), &s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("BUCKET_NAME")),
			Key:    aws.String(hash),
			Body:   descriptor,
		}); err != nil {
			return
		}

		rock.Rock.Set(context.Background(), hash, "1", 0)
	} else if err != nil {
		return
	}

	// Tambah versi baru
	newVersion := model.FileVersion{
		FileID:  fl.ID,
		Version: lastVersion.Version + 1,
		Hash:    hash,
		Size:    fileHeader.Size,
	}

	if err = tx.Create(&newVersion).Error; err != nil {
		return
	}

	tx.Commit()
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "saved",
		"version": fmt.Sprint(newVersion.Version),
	})
}

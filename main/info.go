package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mathrock-xyz/starducc/main/auth"
	"github.com/mathrock-xyz/starducc/main/db"
	"github.com/mathrock-xyz/starducc/main/db/model"
	"gorm.io/gorm"
)

func info(ctx echo.Context) (err error) {
	userid, name := auth.UserId(ctx), ctx.Param("name")
	if name == "" {
		return echo.NewHTTPError(
			http.StatusBadRequest, "file name is required",
		)
	}

	var file model.File

	// Query file dan preload versions yang diurutkan descending (versi terbaru di indeks pertama)
	if err = db.DB.Where("name = ? AND user_id = ?", name, userid).
		Preload("Versions", func(db *gorm.DB) *gorm.DB {
			return db.Order("version DESC").Limit(1)
		}).
		First(&file).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(
				http.StatusNotFound, "file not found",
			)
		}
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"database query failed",
		)
	}

	currentver := 0
	if len(file.Versions) > 0 {
		currentver = file.Versions[0].Ver
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"name":            file.Name,
		"hash":            file.Hash,
		"size":            file.Size,
		"current_version": currentver,
	})
}

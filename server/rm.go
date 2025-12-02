package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mathrock-xyz/starducc/server/auth"
	"github.com/mathrock-xyz/starducc/server/db"
	"github.com/mathrock-xyz/starducc/server/db/model"
)

func rm(ctx echo.Context) (err error) {
	userID := auth.UserId(ctx)
	if userID == "" {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"error": "please login first",
		})
	}

	fileName := ctx.Param("name")
	if fileName == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "file name is required",
		})
	}

	tx := db.DB.Begin()
	defer tx.Rollback()

	result := tx.
		Where("name = ? AND user_id = ? AND locked = ?", fileName, userID, false).
		Delete(&model.File{})

	if result.Error != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "delete failed",
		})
	}

	if result.RowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"error": "file not found",
		})
	}

	tx.Commit()

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "file deleted",
	})
}

package auth

import "github.com/labstack/echo/v4"

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {}
}

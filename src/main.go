package main

import "github.com/labstack/echo/v4"

func main() {
	mux := echo.New()

	mux.Start(":8000")
}

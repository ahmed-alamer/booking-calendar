package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	api := echo.New()
	api.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]string{"System": "OK"})
	})

	api.Logger.Fatal(api.Start(":8080"))
}

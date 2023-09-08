package rest

import (
	"github.com/labstack/echo/v4"
	"log"
)

func Server() {
	e := echo.New()

	e.GET("/system/alive", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "I'm alive!",
		})
	})

	log.Fatalf("Server error: %v", e.Start(":9800"))
}

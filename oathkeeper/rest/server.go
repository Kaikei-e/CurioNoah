package rest

import (
	"ghostkeeper/logger"
	"github.com/labstack/echo/v4"
)

func Server() {
	e := echo.New()

	e.GET("/system/alive", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "I'm alive!",
		})
	})

	err := e.Start(":9800")
	if err != nil {
		logger.Logger.Error("Server error: %v", err)
		panic(err)
	}

}

package server

import (
	"github.com/labstack/echo/v4"
)

func Server() {
	e := echo.New()
	err := e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	if err != nil {
		e.Logger.Errorf("error: %v. maybe serer is down", err)
	}

	apiV1 := e.Group("/api/v1")
	apiV1.Use()
	{
		err := apiV1.GET("/fetch-feeds", func(c echo.Context) error {
			return c.String(200, "Hello, World!")
		})
		if err != nil {
			e.Logger.Errorf("failed to fetch feeds. error: %v.", err)
		}
	}

	e.Logger.Fatal(e.Start(":9000"))
}

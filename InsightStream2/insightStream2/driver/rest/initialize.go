package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"insightStream2/gateway/server"
)

func Initialize() {
	e := echo.New()
	rest := server.NewServerGateway()

	e.Use(middleware.Logger())

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173",
			"http://localhost:4173", "http://127.0.0.1:4173",
			"http://curionoah.com:5173", "http://curionoah.com:4173",
		},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	e.GET("/system/health", func(c echo.Context) error {
		err := rest.HealthCheck()
		if err != nil {
			panic(err)
		}

		return c.String(200, "Keep Alive")
	})

	apiV1 := e.Group("/api/v1")
	apiV1.Use()
	{
		panic("implement me")
	}

	e.Logger.Fatal(e.Start(":9110"))

}

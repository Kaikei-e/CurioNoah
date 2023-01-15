package server

import (
	"feedflare/collector/fetchFeeds"
	"feedflare/collector/testdata"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mmcdole/gofeed"
)

func Server() {
	e := echo.New()
	err := e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	if err != nil {
		e.Logger.Errorf("error: %v. maybe serer is down", err)
	}

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	apiV1 := e.Group("/api/v1")
	apiV1.Use()
	{
		err := apiV1.GET("/fetch-feeds", func(c echo.Context) error {
			feeds, err := fetchFeeds.MultiFeed(testdata.FeedList)
			if err != nil {
				e.Logger.Errorf("error: %v. maybe serer is down", err)
				return err
			}

			var feedsFormatted []gofeed.Feed
			for _, feed := range feeds {
				feedFormatted := *feed
				feedsFormatted = append(feedsFormatted, feedFormatted)
			}

			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")

			return c.JSON(200, feedsFormatted)
		})
		if err != nil {
			e.Logger.Errorf("failed to fetch feeds. error: %v.", err)
		}
	}

	e.Logger.Fatal(e.Start(":9000"))
}

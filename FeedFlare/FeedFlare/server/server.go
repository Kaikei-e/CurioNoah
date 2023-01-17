package server

import (
	"feedflare/collector/testdata"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mmcdole/gofeed"
)

func Server() {
	e := echo.New()
	e.Use(middleware.Logger())

	err := e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	if err != nil {
		e.Logger.Errorf("error: %v. maybe serer is down", err)
	}

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
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

		fetchFeed := apiV1.Group("/fetch-feed")
		fetchFeed.Use()
		{
			err := apiV1.GET("/stored-all", func(c echo.Context) error {
				e.Logger.Info("stored-all api is called")

				feeds, err := fetchFeed.MultiFeed(testdata.FeedList)
				if err != nil {
					e.Logger.Errorf("error: %v. maybe serer is down", err)
					return err
				}

				e.Logger.Infof("feeds were fetched: feed number is %v", len(feeds))

				var feedsFormatted []gofeed.Feed
				for _, feed := range feeds {
					feedFormatted := *feed
					feedsFormatted = append(feedsFormatted, feedFormatted)
				}

				c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
				c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
				c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")

				e.Logger.Info("response header is set")

				return c.JSON(200, feedsFormatted)
			})
			if err != nil {
				e.Logger.Errorf("failed to fetch feeds. error: %v.", err)
			}
		}

		registerFeed := apiV1.Group("/registerFeed-feed")
		registerFeed.Use()
		{
			err := apiV1.POST("/registerFeed", func(c echo.Context) error {
				e.Logger.Info("registerFeed api is called")

				return c.String(200, "registerFeed api is called")

			})
			if err != nil {
				e.Logger.Errorf("failed to registerFeed feed. error: %v.", err)

			}
		}
	}

	e.Logger.Fatal(e.Start(":9000"))
}

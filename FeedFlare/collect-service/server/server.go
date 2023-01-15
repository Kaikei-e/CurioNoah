package server

import (
	"feedflare/collector/fetchFeeds"
	"feedflare/collector/testdata"
	"github.com/labstack/echo/v4"
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

	apiV1 := e.Group("/api/v1")
	apiV1.Use()
	{
		err := apiV1.GET("/fetch-feeds", func(c echo.Context) error {
			feeds, err := fetchFeeds.MultiFeed(testdata.FeedList)
			if err != nil {
				e.Logger.Errorf("error: %v. maybe serer is down", err)
				return err
			}

			var feedsFromatted []gofeed.Feed
			for _, feed := range feeds {
				feedFormatted := *feed
				feedsFromatted = append(feedsFromatted, feedFormatted)
			}

			return c.JSON(200, feedsFromatted)
		})
		if err != nil {
			e.Logger.Errorf("failed to fetch feeds. error: %v.", err)
		}
	}

	e.Logger.Fatal(e.Start(":9000"))
}
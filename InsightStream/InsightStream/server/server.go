package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mmcdole/gofeed"
	"insightstream/adaptor"
	register "insightstream/collector/registerFeed"
	"insightstream/ent"
	"insightstream/repository/readfeed"
	"insightstream/restorerss"
	"sort"
)

func Server(cl *ent.Client) {
	e := echo.New()
	e.Use(middleware.Logger())

	err := e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	if err != nil {
		e.Logger.Errorf("error: %v. maybe serer is down", err)
	}

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

	apiV1 := e.Group("/api/v1")
	apiV1.Use()
	{

		fetchFeed := apiV1.Group("/fetch-feed")
		fetchFeed.Use()
		{
			// TODO will rotate the logics to small functions
			err := fetchFeed.GET("/stored-all", func(c echo.Context) error {
				e.Logger.Info("stored-all api is called")

				//feeds, err := fetchFeeds.ParallelizeFetch(testdata.FeedList)
				//feedEnt, err := readfeed.QueryByTen(cl)
				feedEnt, err := readfeed.QueryByTwenty(cl)
				if err != nil {
					e.Logger.Errorf("error: %v. maybe sever is down", err)
					// TODO FIX: return error
					err := c.JSON(500, err)
					if err != nil {
						e.Logger.Errorf("error: %v. maybe sever is down", err)

					}
				}

				feeds, err := restorerss.FeedExchange(feedEnt)
				if err != nil {
					e.Logger.Errorf("error: %v. maybe sever is down", err)
					// TODO FIX: return error
					err := c.JSON(500, err)
					if err != nil {
						e.Logger.Errorf("error: %v. maybe sever is down", err)

					}
				}

				e.Logger.Infof("feeds were fetched: feed number is %v", len(feeds))

				var feedsFormatted []gofeed.Feed
				for _, feed := range feeds {
					feedFormatted := *feed
					feedsFormatted = append(feedsFormatted, feedFormatted)
				}

				sort.Slice(feedsFormatted, func(i, j int) bool {
					return feedsFormatted[i].UpdatedParsed.After(*feedsFormatted[j].UpdatedParsed)
				})

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

		registerFeed := apiV1.Group("/register-feed")
		registerFeed.Use()
		{
			register.RegisterHandler(registerFeed, cl)

		}
	}

	// change the architecture to onion architecture like
	apiV2 := e.Group("/api/v2")
	apiV2.Use()
	{
		fetchFeed := apiV2.Group("/fetch-feed")
		fetchFeed.Use()
		{
			err := fetchFeed.GET("/stored-all", func(c echo.Context) error {
				err := adaptor.CollectAll(c, cl)

				return err
			})

			if err != nil {
				e.Logger.Errorf("failed to fetch feeds. error: %v.", err)

			}

		}
	}

	e.Logger.Fatal(e.Start(":9000"))
}

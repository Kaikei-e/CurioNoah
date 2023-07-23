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
	"insightstream/restorerss/manageFeedsAmount"
	"log"
	"sort"
	"strconv"
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
			fetchFeed.GET("/stored-all", func(c echo.Context) error {
				e.Logger.Info("stored-all api is called")

				qpStr := c.QueryParam("page")
				qp, err := strconv.Atoi(qpStr)
				if err != nil {
					e.Logger.Errorf("failed to convert string to int. error: %v", err)
				}

				// the twenty is the number of list to send the followList to Front
				feedEnt, hadExceeded, err := readfeed.QueryByTwenty(cl, qp)
				if err != nil {
					e.Logger.Errorf("error: %v. maybe sever is down", err)
					// TODO FIX: return error
					err := c.JSON(500, err)
					if err != nil {
						e.Logger.Errorf("error: %v. maybe sever is down", err)
					}
				}

				c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
				c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
				c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")

				if feedEnt == nil && hadExceeded {
					emptyRes := Response{
						Feeds:       nil,
						HadExceeded: hadExceeded,
					}

					e.Logger.Info("no more feeds to fetch")
					return c.JSON(200, emptyRes)
				}

				feeds, err := restorerss.EntFollowListExchangeToGofeed(feedEnt)
				if err != nil {
					e.Logger.Errorf("error: %v. maybe sever is down", err)
					// TODO FIX: return error
					err := c.JSON(500, err)
					if err != nil {
						e.Logger.Errorf("error: %v. maybe sever is down", err)
					}
				}

				reducedFeeds, err := manageFeedsAmount.ReduceToLatestThreeItems(feeds)
				if err != nil {
					e.Logger.Errorf("failed to reduce baseFeeds. error: %v", err)
				}

				e.Logger.Infof("feeds were fetched: feed number is %v", len(reducedFeeds))

				var feedsFormatted []gofeed.Feed
				for _, feed := range reducedFeeds {
					feedFormatted := feed
					feedsFormatted = append(feedsFormatted, feedFormatted)
				}

				sort.Slice(feedsFormatted, func(i, j int) bool {
					return feedsFormatted[i].UpdatedParsed.After(*feedsFormatted[j].UpdatedParsed)
				})

				e.Logger.Info("response header is set")

				res := Response{
					Feeds:       feedsFormatted,
					HadExceeded: hadExceeded,
				}

				return c.JSON(200, res)
			})
		}

		registerFeed := apiV1.Group("/register-feed")
		registerFeed.Use()
		{
			register.RegisterHandler(registerFeed, cl)
		}

		infiniteScroll := apiV1.Group("/infinite-scroll")
		infiniteScroll.Use()
		{
			infiniteScroll.GET("/stored-all", func(c echo.Context) error {
				err := adaptor.InfiniteFetching(c, cl)
				return err
			})
		}

		search := apiV1.Group("/search")
		search.Use()
		{
			search.GET("/baseFeeds", func(c echo.Context) error {
				searchedFeeds, err := adaptor.SearchFeeds(c, cl)
				if err != nil {
					log.Fatalln(err)
				}

				return c.JSON(200, searchedFeeds)
			})
		}

	}

	// change the architecture to onion architecture like
	apiV2 := e.Group("/api/v2")
	apiV2.Use()
	{
		fetchFeed := apiV2.Group("/fetch-feed")
		fetchFeed.Use()
		{
			fetchFeed.GET("/stored-all", func(c echo.Context) error {
				err := adaptor.CollectAll(c, cl)
				return err
			})
		}
	}

	e.Logger.Fatal(e.Start(":9000"))
}

type Response struct {
	Feeds       []gofeed.Feed `json:"feeds"`
	HadExceeded bool          `json:"hadExceeded"`
}

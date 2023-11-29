package registerFeed

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"insightstream/collector/fetchFeedDomain"
	"insightstream/ent"
	"insightstream/models/apiexcahnge"
)

// TODO will implement unit tests
func RegisterHandler(g *echo.Group, cl *ent.Client) {
	register := g.Group("/store")
	register.Use()
	{
		register.POST("/single", func(c echo.Context) error {
			c.Logger().Info("single api is called")

			sf := new(apiexcahnge.FeedRegister)

			if err := c.Bind(sf); err != nil {
				c.Logger().Errorf("error in bind: %v", err)
				return c.JSON(400, errors.New("invalid request"))
			}

			c.Logger().Infof("single feed registering started: %v", sf)

			feeds, err := fetchFeedDomain.MultiFeed([]string{sf.URL})
			if err != nil {
				c.Logger().Errorf("error in fetch feed: %v", err)
				return errors.New(fmt.Sprintf("fetch %s: %v", sf.URL, err))
			}

			err = RegisterSingle(sf.URL, feeds[0], cl)
			if err != nil {
				c.Logger().Errorf("error in register feed: %v", err)
				return errors.New(fmt.Sprintf("register %s: %v", sf.URL, err))
			}

			c.Logger().Infof("single feed is registered: %v", sf)

			response := struct {
				Message   string `json:"message"`
				TargetURL string `json:"target_url"`
			}{
				Message:   "success",
				TargetURL: sf.URL,
			}

			return c.JSON(200, response)
		})

	}
}

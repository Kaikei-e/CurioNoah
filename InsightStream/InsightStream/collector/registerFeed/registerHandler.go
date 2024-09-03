package registerFeed

import (
	"errors"
	"fmt"
	"insightstream/collector/fetchFeedDomain"
	"insightstream/ent"
	"insightstream/models/apiexcahnge"
	"net/url"

	"github.com/labstack/echo/v4"
)

// TODO will implement unit tests
func RegisterHandler(g *echo.Group, cl *ent.Client) {
	register := g.Group("/store")
	register.Use()
	{
		register.POST("/single", func(c echo.Context) error {
			c.Logger().Info("single url registering api is called")

			sf := new(apiexcahnge.FeedRegister)

			if err := c.Bind(sf); err != nil {
				c.Logger().Errorf("error in bind: %v", err)
				return c.JSON(400, errors.New("invalid request"))
			}

			parsedURL, err := url.Parse(sf.URL)
			if err != nil {
				c.Logger().Errorf("error in parse: %v", err)
				return c.JSON(400, errors.New("invalid request"))
			}

			parsedURL.RawQuery = ""

			c.Logger().Infof("single feed registering started: %v", sf)

			feeds, err := fetchFeedDomain.MultiFeed([]string{parsedURL.String()})
			if err != nil {
				c.Logger().Errorf("error in fetch feed: %v", err)
				return fmt.Errorf("fetch %s: %v", parsedURL, err)
			}

			err = RegisterSingle(parsedURL.String(), feeds[0], cl)
			if err != nil {
				c.Logger().Errorf("error in register feed: %v", err)
				return fmt.Errorf("register %s: %v", parsedURL, err)
			}

			c.Logger().Infof("single feed is registered: %v", sf)

			response := struct {
				Message   string `json:"message"`
				TargetURL string `json:"target_url"`
			}{
				Message:   "success",
				TargetURL: parsedURL.String(),
			}

			return c.JSON(200, response)
		})

	}
}

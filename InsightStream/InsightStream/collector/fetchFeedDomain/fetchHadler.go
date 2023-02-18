package fetchFeedDomain

import (
	"github.com/labstack/echo/v4"
	"insightstream/ent"
)

func FetchHandler(eg *echo.Group, cl *ent.Client) {
	stored := eg.Group("/stored")
	{
		stored.GET("/all", func(c echo.Context) error {
			return nil
		})

	}

	realtime := eg.Group("/realtime")
	{
		realtime.GET("/latest", func(c echo.Context) error {

			return nil
		})
	}

}

package adaptor

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"insightstream/ent"
)

func SearchFeeds(c echo.Context, cl *ent.Client) error {

	queryParams := c.QueryParams()
	keyword := queryParams.Get("searchWord")

	fmt.Println("keyword: ", keyword)

	return nil
}

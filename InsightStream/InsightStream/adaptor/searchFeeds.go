package adaptor

import (
	"github.com/labstack/echo/v4"
	"insightstream/domain/searchWord"
	"insightstream/ent"
	"insightstream/models/feeds"
	"insightstream/usecase/searchUsecase"
)

func SearchFeeds(c echo.Context, cl *ent.Client) ([]feeds.EachFeed, error) {

	queryParams := c.QueryParams()
	title := queryParams.Get("title")
	description := queryParams.Get("description")

	keyword := searchWord.SearchWord{
		Title:       title,
		Description: description,
	}

	titleOrDescription, err := searchUsecase.SearchByTitleOrDescription(keyword, cl)
	if err != nil {
		return nil, err
	}

	return titleOrDescription, nil
}

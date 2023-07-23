package adaptor

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"insightstream/domain/searchWord"
	"insightstream/domain/searchedFeed"
	"insightstream/driver/parser"
	"insightstream/ent"
	"insightstream/usecase/searchUsecase"
)

func SearchFeeds(c echo.Context, cl *ent.Client) ([]searchedFeed.ByTitleOrDescription, error) {

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

	var formattedFeeds []searchedFeed.ByTitleOrDescription
	for _, tod := range titleOrDescription {
		description, err := parser.HTMLToDoc(tod.Description)
		if err != nil {
			err := fmt.Errorf("failed to parse html: %v", err)
			fmt.Println(err)
			continue
		}

		tod.Description = description
		formattedFeeds = append(formattedFeeds, tod)
	}

	return formattedFeeds, nil
}

package adaptor

import (
	"fmt"
	"insightstream/domain/searchWord"
	"insightstream/domain/searchedFeed"
	"insightstream/driver/parser"
	"insightstream/ent"
	"insightstream/usecase/searchUsecase"
	"log/slog"

	"github.com/labstack/echo/v4"
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
		slog.Error("failed to search by title or description: %v", err)
		return nil, fmt.Errorf("failed to search by title or description: %w", err)
	}

	var formattedFeeds []searchedFeed.ByTitleOrDescription
	for _, tod := range titleOrDescription {
		description, err := parser.HTMLToDoc(tod.Description)
		if err != nil {
			err := fmt.Errorf("failed to parse html: %v", err)
			slog.Error(err.Error())
			continue
		}

		tod.Description = description
		formattedFeeds = append(formattedFeeds, tod)
	}

	return formattedFeeds, nil
}

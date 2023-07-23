package searchUsecase

import (
	"context"
	"insightstream/domain/searchWord"
	"insightstream/domain/searchedFeed"
	"insightstream/ent"
	"insightstream/port/searchPort"
	"insightstream/restorerss/convertToResponse"
)

func SearchByTitleOrDescription(searchWord searchWord.SearchWord, cl *ent.Client) ([]searchedFeed.ByTitleOrDescription, error) {
	ctx := context.Background()

	searchImpl := searchPort.Impl{}
	searchedFeeds, err := searchImpl.SearchByTitleOrDescription(searchWord, cl, ctx)
	if err != nil {
		return nil, err
	}

	toSearchWord, err := convertToResponse.ToSearchWord(searchedFeeds)
	if err != nil {
		return nil, err
	}

	return toSearchWord, nil
}

package searchUsecase

import (
	"context"
	"insightstream/domain/searchWord"
	"insightstream/ent"
	"insightstream/models/feeds"
	"insightstream/port/searchPort"
	"insightstream/restorerss"
)

func SearchByTitleOrDescription(searchWord searchWord.SearchWord, cl *ent.Client) ([]feeds.EachFeed, error) {
	ctx := context.Background()

	searchImpl := searchPort.Impl{}
	searchedFeeds, err := searchImpl.SearchByTitleOrDescription(searchWord, cl, ctx)
	if err != nil {
		return nil, err
	}

	gfs, err := restorerss.ExchangeEntFeedsToGofeeds(searchedFeeds)
	if err != nil {
		return nil, err
	}
	return gfs, nil
}

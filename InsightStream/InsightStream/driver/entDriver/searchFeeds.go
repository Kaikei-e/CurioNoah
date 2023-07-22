package entDriver

import (
	"context"
	"insightstream/domain/searchWord"
	"insightstream/ent"
	entFeeds "insightstream/ent/feeds"
)

func SearchFeeds(searchWord searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error) {
	const limit = 40

	feeds, err := cl.Feeds.Query().Select().Limit(limit).Where(
		entFeeds.Or(
			entFeeds.TitleContains(searchWord.Title),
			entFeeds.DescriptionContains(searchWord.Description),
		),
	).Order(
		ent.Desc("dt_updated"),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}

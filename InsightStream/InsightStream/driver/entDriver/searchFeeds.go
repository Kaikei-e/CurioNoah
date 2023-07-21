package entDriver

import (
	"context"
	"insightstream/domain/searchWord"
	"insightstream/ent"
	entFeeds "insightstream/ent/feeds"
)

func SearchFeeds(searchWord searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error) {
	const limit = 100

	feeds, err := cl.Feeds.Query().Select("title", "description").Limit(20).Where(
		entFeeds.Or(
			entFeeds.TitleContains(searchWord.Title),
			entFeeds.DescriptionContains(searchWord.Description),
		),
	).Order(
		ent.Desc("dt_updated"),
	).Offset(limit).All(ctx)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}

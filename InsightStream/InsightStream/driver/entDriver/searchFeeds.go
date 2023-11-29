package entDriver

import (
	"context"
	"insightstream/ent"
	entFeeds "insightstream/ent/feeds"
)

func SearchFeeds(keyword string, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error) {
	if keyword == "" {
		return []*ent.Feeds{}, nil
	}

	const limit = 100

	feeds, err := cl.Feeds.Query().
		Limit(limit).
		Where(
			entFeeds.Or(
				entFeeds.TitleContains(keyword),
				entFeeds.DescriptionContains(keyword),
			),
		).
		Order(ent.Desc("dt_updated")).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return feeds, nil
}

package entDriver

import (
	"context"
	"fmt"
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
		return nil, fmt.Errorf("failed to query feeds: %w", err)
	}

	return feeds, nil
}

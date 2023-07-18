package entDriver

import (
	"context"
	"insightstream/domain/searchWord"
	"insightstream/ent"
)

func SearchFeeds(searchWord searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error) {
	feeds, err := cl.Feeds.Query().Select("title", "description").Limit(100).Where().All(ctx)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}

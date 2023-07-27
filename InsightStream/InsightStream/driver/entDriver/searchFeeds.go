package entDriver

import (
	"context"
	"errors"
	"fmt"
	"insightstream/domain/searchWord"
	"insightstream/ent"
	entFeeds "insightstream/ent/feeds"
)

func SearchFeeds(searchWord searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error) {
	// This is the limit of the number of feeds to be displayed on the screen.
	// not intentional number, just for now.
	const limit = 100

	feeds, err := cl.Feeds.Query().Select().Limit(limit).Where(
		entFeeds.Or(
			entFeeds.TitleContains(searchWord.Title),
			entFeeds.DescriptionContains(searchWord.Description),
		),
	).Order(
		ent.Desc("dt_updated"),
	).All(ctx)
	if err != nil {
		return nil, errors.New("failed to query: " + err.Error())
	}

	fmt.Println("baseFeeds length is : ", len(feeds))

	return feeds, nil
}

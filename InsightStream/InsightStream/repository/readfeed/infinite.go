package readfeed

import (
	"context"
	"fmt"
	"insightstream/ent"
)

func InfiniteScroll(cl *ent.Client, queryParam int) ([]*ent.Feeds, bool, error) {
	var hadExceeded bool

	ctx := context.Background()

	const limit = 40
	offset := limit * (queryParam - 1)

	rowAmount, err := CountFeeds()
	if err != nil {
		return nil, false, err
	}

	if rowAmount < limit {
		hadExceeded = true
	} else {
		hadExceeded = false
	}

	feeds, err := cl.Feeds.Query().
		Order(ent.Desc("dt_updated")).
		Order(ent.Desc("feed_url")).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("failed to query by ten: %v", err)
	}

	return feeds, hadExceeded, nil

}

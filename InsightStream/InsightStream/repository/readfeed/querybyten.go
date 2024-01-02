package readfeed

import (
	"context"
	"fmt"
	"insightstream/ent"
)

func QueryByTen(cl *ent.Client) ([]*ent.FollowLists, error) {
	ctx := context.Background()

	fl, err := cl.FollowLists.Query().
		Order(ent.Desc("id")).
		Limit(10).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to query by ten: %v", err)
	}

	return fl, nil

}

func QueryByTwenty(cl *ent.Client, queryParam int) ([]*ent.FollowLists, bool, error) {
	var hadExceeded bool

	ctx := context.Background()

	const limit = 20
	offset := limit * (queryParam - 1)

	rowAmount, err := CountFollowList()
	if err != nil {
		return nil, false, fmt.Errorf("failed to count followlist: %v", err)
	}

	if rowAmount < limit {
		hadExceeded = true
	} else {
		hadExceeded = false
	}

	fl, err := cl.FollowLists.Query().
		Order(ent.Desc("dt_updated")).
		Order(ent.Desc("id")).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, false, fmt.Errorf("failed to query by %v: %v", limit, err)
	}

	return fl, hadExceeded, nil
}

func trimFeedList(fl []*ent.FollowLists, amount int) []*ent.FollowLists {
	const theLimit = 20

	// if fl is nil or empty, return an empty slice
	if fl == nil || len(fl) == 0 {
		return []*ent.FollowLists{}
	}

	// if amount is less than or equal to theLimit, return the full array
	if amount <= theLimit {
		return fl
	}

	// If amount exceeds the limit, trim the list to contain only the last 'theLimit' items
	startIndex := amount - theLimit
	return fl[startIndex:]
}

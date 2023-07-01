package readfeed

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent"
)

func QueryByTen(cl *ent.Client) ([]*ent.FollowList, error) {
	ctx := context.Background()

	fl, err := cl.FollowList.Query().
		Order(ent.Desc("id")).
		Limit(10).
		All(ctx)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to query by ten: %v", err))
	}

	return fl, nil

}

func QueryByTwenty(cl *ent.Client, queryParam int) ([]*ent.FollowList, bool, error) {
	var hadExceeded bool

	ctx := context.Background()

	const limit = 20
	offset := limit * (queryParam - 1)

	rowAmount, err := CountFollowList()
	if err != nil {
		return nil, false, errors.New(fmt.Sprintf("failed to count the row: %v", err))
	}

	if rowAmount < limit {
		hadExceeded = true
	} else {
		hadExceeded = false
	}

	fl, err := cl.FollowList.Query().
		Order(ent.Desc("dt_updated")).
		Order(ent.Desc("id")).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, false, errors.New(fmt.Sprintf("failed to query by ten: %v", err))
	}

	return fl, hadExceeded, nil
}

func trimFeedList(fl []*ent.FollowList, amount int) []*ent.FollowList {
	const theLimit = 20

	// if fl is nil or empty, return an empty slice
	if fl == nil || len(fl) == 0 {
		return []*ent.FollowList{}
	}

	// if amount is less than or equal to theLimit, return the full array
	if amount <= theLimit {
		return fl
	}

	// If amount exceeds the limit, trim the list to contain only the last 'theLimit' items
	startIndex := amount - theLimit
	return fl[startIndex:]
}

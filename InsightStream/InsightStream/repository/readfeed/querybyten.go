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

func QueryByTwenty(cl *ent.Client, queryParam int) ([]*ent.FollowList, error) {
	ctx := context.Background()

	const limit = 20
	var limitAmount int
	if queryParam > 0 {
		limitAmount = queryParam * 20
	} else {
		limitAmount = limit
	}

	fl, err := cl.FollowList.Query().
		Order(ent.Desc("dt_updated")).
		Limit(limitAmount).
		All(ctx)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to query by ten: %v", err))
	}

	return fl, nil
}

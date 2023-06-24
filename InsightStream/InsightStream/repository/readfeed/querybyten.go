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

	limit := 20
	if queryParam > 0 {
		limit = queryParam + limit
	}

	fl, err := cl.FollowList.Query().
		Order(ent.Desc("dt_updated")).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to query by ten: %v", err))
	}

	return fl, nil
}

package readfeed

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent"
	"strconv"
)

func QueryByTen(cl *ent.Client) ([]*ent.FollowList, error) {
	var following ent.FollowList
	ctx := context.Background()

	fl, err := cl.FollowList.Query().
		Order(ent.Desc(strconv.Itoa(following.ID))).
		Limit(10).
		All(ctx)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to query by ten: %v", err))
	}

	return fl, nil

}

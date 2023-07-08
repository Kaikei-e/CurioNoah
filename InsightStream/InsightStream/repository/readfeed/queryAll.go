package readfeed

import (
	"context"
	"fmt"
	"insightstream/ent"
	"insightstream/ent/followlist"
)

func QueryAll(cl *ent.Client) ([]*ent.FollowList, error) {
	fmt.Println("query all")
	ctx := context.Background()

	fl, err := cl.FollowList.Query().
		Select(followlist.FieldID, followlist.FieldURL,
			followlist.FieldLink, followlist.FieldLinks,
			followlist.FieldDtLastInserted).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to query all: %v", err)
	}

	fmt.Println("query all done: ", len(fl), "rows")

	return fl, nil
}

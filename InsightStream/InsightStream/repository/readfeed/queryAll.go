package readfeed

import (
	"context"
	"fmt"
	"insightstream/ent"
	"insightstream/ent/followlists"
)

func QueryAll(cl *ent.Client) ([]*ent.FollowLists, error) {
	fmt.Println("query all")
	ctx := context.Background()

	fl, err := cl.FollowLists.Query().
		Select(followlists.FieldID, followlists.FieldURL,
			followlists.FieldLink, followlists.FieldLinks,
			followlists.FieldDtLastInserted).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to query all: %v", err)
	}

	fmt.Println("query all done: ", len(fl), "rows")

	return fl, nil
}

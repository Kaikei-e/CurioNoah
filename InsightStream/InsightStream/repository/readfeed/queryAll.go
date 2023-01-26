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

	fl := cl.FollowList.Query().
		Select(followlist.FieldID, followlist.FieldURL,
			followlist.FieldLink, followlist.FieldLinks, followlist.FieldDtLastInserted).
		AllX(ctx)

	return fl, nil
}

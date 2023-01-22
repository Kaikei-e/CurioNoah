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

	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("failed to query all data: %v", err))
	//}

	return fl, nil
}

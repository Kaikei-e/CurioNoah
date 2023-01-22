package registerFeed

import (
	"context"
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"insightstream/ent"
	"insightstream/ent/followlist"
	"insightstream/models/feeds"
	"time"
)

func Update(fds []*gofeed.Feed, cl *ent.Client) error {
	ctx := context.Background()

	n := time.Now()

	// TODO this updating method is weak. also update one by one is not good.
	for _, fd := range fds {
		var ls feeds.FeedLink
		ls.Link = []string{fd.Link}

		_, err := cl.FollowList.Update().
			Where(
				followlist.IsActive(true),
				followlist.URL(fd.Link), //caution this url means site link without feed link
			).
			SetDtLastInserted(n). // is this necessary?
			SetDtUpdated(n).
			SetLink(fd.FeedLink).
			SetURL(fd.Link).
			SetLinks(ls).
			Save(ctx)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to update: %v", err))
		}
	}

	return nil
}

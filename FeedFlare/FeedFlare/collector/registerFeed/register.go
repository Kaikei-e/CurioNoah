package registerFeed

import (
	"feedflare/ent"
	"github.com/mmcdole/gofeed"
)

func Register(feeds []*gofeed.Feed) error {
	create := ent.FollowListCreate{}

	return nil
}

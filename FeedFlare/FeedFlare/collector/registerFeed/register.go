package registerFeed

import (
	"context"
	"feedflare/ent"
	"feedflare/repository"
	"github.com/mmcdole/gofeed"
)

func Register(feeds []*gofeed.Feed) error {
	client := repository.InitConnection()

	ctx := context.Background()
	bulk := make([]*ent.FollowListCreate, 0, len(feeds))

	for i, feed := range feeds {
		bulk[i] = client.FollowList.Create().
			SetTitle(feed.Title).
			SetURL(feed.Link).
			SetDescription(feed.Description).
			SetLanguage(feed.Language).
			SetIsActive(true).
			SetIsFavorite(false).
			SetIsRead(false).
			SetIsUpdated(false).
			SetLink(feed.Link).
			SetLinks(feed.Links).


	}

	return nil
}

package registerFeed

import (
	"context"
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"insightstream/ent"
	"insightstream/repository"
)

func RegisterSingle(feed *gofeed.Feed) error {
	client := repository.InitConnection()

	ctx := context.Background()

	_, err := client.FollowList.Create().
		SetTitle(feed.Title).
		SetURL(feed.Link).
		SetDescription(feed.Description).
		SetLanguage(feed.Language).
		SetIsActive(true).
		SetIsFavorite(false).
		SetIsRead(false).
		SetIsUpdated(false).
		SetLink(feed.Link).Save(ctx)

	if err != nil {
		return errors.New(fmt.Sprintf("failed to register feed: %v", err))
	}

	return nil
}

func RegisterMulti(feeds []*gofeed.Feed) error {
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
			SetLink(feed.Link)
		//SetLinks(feed.Links)

	}

	if _, err := client.FollowList.CreateBulk(bulk...).Save(ctx); err != nil {
		return errors.New(fmt.Sprintf("failed to register feed: %v", err))
	}

	return nil
}

package registerFeed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"insightstream/ent"
	"insightstream/repository"
)

// TODO will implement unit tests
func RegisterSingle(feed *gofeed.Feed, cl *ent.Client) error {
	//client := repository.InitConnection()

	var links []string
	for _, item := range feed.Items {
		for _, link := range item.Links {
			links = append(links, link)
		}
	}

	flattenLinks, err := json.Marshal(links)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to marshal links: %v", err))
	}

	ctx := context.Background()

	_, err = cl.FollowList.Create().
		SetTitle(feed.Title).
		SetURL(feed.Link).
		SetDescription(feed.Description).
		SetLanguage(feed.Language).
		SetIsActive(true).
		SetIsFavorite(false).
		SetIsRead(false).
		SetIsUpdated(false).
		SetLink(feed.Link).
		SetLinks(string(flattenLinks)).
		Save(ctx)

	if err != nil {
		return errors.New(fmt.Sprintf("failed to register feed: %v", err))
	}

	return nil
}

// TODO will implement unit tests
func RegisterMulti(feeds []*gofeed.Feed, cl *ent.Client) error {
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

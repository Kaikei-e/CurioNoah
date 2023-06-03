package registerFeed

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"insightstream/ent"
	"insightstream/models/feeds"
)

// TODO will implement unit tests
func RegisterSingle(feed *gofeed.Feed, cl *ent.Client) error {
	//client := repository.InitConnection()

	var links []string
	links = append(links, feed.Links...)

	var linksJson feeds.FeedLink
	linksJson.Link = links

	var fis []feeds.FeedItem
	for _, item := range feed.Items {
		var authors []string
		for _, author := range item.Authors {
			authors = append(authors, author.Name)
		}

		fis = append(fis, feeds.FeedItem{
			Id:              uuid.New(),
			ItemDescription: item.Title,
			ItemLink:        item.Link,
			ItemTitle:       item.Title,
			Updated:         item.Updated,
			UpdatedParsed:   item.UpdatedParsed,
			Published:       item.Published,
			PublishedParsed: item.PublishedParsed,
			Authors:         authors,
			GUID:            item.GUID,
			Categories:      item.Categories,
		})
	}

	ctx := context.Background()

	_, err := cl.FollowList.Create().
		SetTitle(feed.Title).
		SetURL(feed.Link).
		SetDescription(feed.Description).
		SetLanguage(feed.Language).
		SetItemDescription(fis).
		SetIsActive(true).
		SetIsFavorite(false).
		SetIsRead(false).
		SetIsUpdated(false).
		SetLink(feed.FeedLink).
		SetLinks(linksJson).
		Save(ctx)

	if err != nil {
		return errors.New(fmt.Sprintf("failed to register feed: %v", err))
	}

	return nil
}

// TODO will implement unit tests
func RegisterMulti(feeds []*gofeed.Feed, cl *ent.Client) error {
	ctx := context.Background()
	bulk := make([]*ent.FollowListCreate, 0, len(feeds))

	for i, feed := range feeds {
		bulk[i] = cl.FollowList.Create().
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

	if _, err := cl.FollowList.CreateBulk(bulk...).Save(ctx); err != nil {
		return errors.New(fmt.Sprintf("failed to register feed: %v", err))
	}

	return nil
}

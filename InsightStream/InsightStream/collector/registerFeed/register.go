package registerFeed

import (
	"context"
	"fmt"
	"insightstream/domain/baseFeeds"
	"insightstream/ent"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

// TODO will implement unit tests
func RegisterSingle(inputLink string, feed *gofeed.Feed, cl *ent.Client) error {
	//client := repository.InitConnection()

	var links []string
	links = append(links, feed.Links...)

	var linksJson baseFeeds.FeedLink
	linksJson.Link = links

	var fis []baseFeeds.FeedItem
	tx, err := cl.Tx(context.Background())
	if err != nil {
		return fmt.Errorf("failed to create transaction: %v", err)
	}

	for _, item := range feed.Items {
		var authors []string
		for _, author := range item.Authors {
			authors = append(authors, author.Name)
		}

		fis = append(fis, baseFeeds.FeedItem{
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
	err = cl.FollowLists.Create().
		SetTitle(feed.Title).
		SetURL(feed.Link).
		SetDescription(feed.Description).
		SetLanguage(feed.Language).
		SetItemDescription(fis).
		SetIsActive(true).
		SetIsFavorite(false).
		SetIsRead(false).
		SetIsUpdated(false).
		SetLink(inputLink).
		SetLinks(linksJson).
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("failed to register feed. URL was conflicted!! : %v", err)
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", commitErr)
	}

	return nil
}

// TODO will implement unit tests
func RegisterMulti(feeds []*gofeed.Feed, cl *ent.Client) error {
	ctx := context.Background()
	builders := make([]*ent.FollowListsCreate, len(feeds))

	err := cl.FollowLists.CreateBulk(builders...).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to register feeds: %v", err)
	}

	return nil
}

package registerFeed

import (
	"context"
	"fmt"
	"insightstream/domain/baseFeeds"
	"insightstream/ent"
	"log"
	"time"
)

func Update(fds []*ent.FollowLists, cl *ent.Client) error {
	ctx := context.Background()
	n := time.Now()

	// TODO this updating method is weak. also update one by one is not good.
	// need to consider how to update all at once.
	for _, fd := range fds {

		var feedItems []baseFeeds.FeedItem
		for _, item := range fd.ItemDescription {
			feedItems = append(feedItems, baseFeeds.FeedItem{
				ItemTitle:       item.ItemTitle,
				ItemLink:        item.ItemLink,
				ItemDescription: item.ItemDescription,
				Updated:         item.Updated,
				UpdatedParsed:   item.UpdatedParsed,
				Published:       item.Published,
				PublishedParsed: item.PublishedParsed,
				GUID:            item.GUID,
				Categories:      item.Categories,
			})
		}

		var links []string
		for _, item := range feedItems {
			links = append(links, item.ItemLink)
		}

		var linksJson baseFeeds.FeedLink
		linksJson.Link = links

		tx, err := cl.Tx(ctx)
		if err != nil {
			return fmt.Errorf("failed to create transaction: %v", err)
		}

		_, err = tx.FollowLists.UpdateOneID(fd.ID).
			SetDtLastInserted(n). // is this necessary?
			SetDtUpdated(n).
			SetLink(fd.Link).
			SetURL(fd.URL).
			SetLinks(linksJson).
			SetItemDescription(feedItems).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to update: %v, feed id is %v", err, fd.ID)
		}

		commitErr := tx.Commit()
		if commitErr != nil {
			log.Printf("failed to commit transaction: %v", commitErr)
			tx.Rollback()
			return fmt.Errorf("failed to commit transaction: %v", commitErr)
		}

		if commitErr != nil {
			return commitErr
		}

	}

	return nil
}

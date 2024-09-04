package registerFeed

import (
	"context"
	"fmt"
	"insightstream/domain/baseFeeds"
	"insightstream/ent"
	"log"
	"log/slog"
	"time"
)

func Update(fds []*ent.FollowLists, cl *ent.Client) error {
	ctx := context.Background()
	n := time.Now()

	tx, err := cl.Tx(ctx)

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

		_, err = tx.FollowLists.UpdateOneID(fd.ID).
			SetDtLastInserted(n). // is this necessary?
			SetDtUpdated(n).
			SetLink(fd.Link).
			SetURL(fd.URL).
			SetLinks(linksJson).
			SetItemDescription(feedItems).
			Save(ctx)
		if err != nil {
			slog.Error("failed to update feed:", "error", err)
			continue
		}
	}

	commitErr := tx.Commit()
	if commitErr != nil {
		log.Printf("failed to commit transaction: %v", commitErr)
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", commitErr)
	}

	return nil
}

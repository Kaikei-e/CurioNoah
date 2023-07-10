package registerFeed

import (
	"context"
	"errors"
	"fmt"
	"insightstream/ent"
	"insightstream/models/feeds"
	"time"
)

func Update(fds []*ent.FollowList, cl *ent.Client) error {
	ctx := context.Background()
	n := time.Now()

	// TODO this updating method is weak. also update one by one is not good.
	// need to consider how to update all at once.
	for _, fd := range fds {

		var feedItems []feeds.FeedItem
		for _, item := range fd.ItemDescription {
			feedItems = append(feedItems, feeds.FeedItem{
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

		var linksJson feeds.FeedLink
		linksJson.Link = links

		_, err := cl.FollowList.UpdateOneID(fd.ID).
			SetDtLastInserted(n). // is this necessary?
			SetDtUpdated(n).
			SetLink(fd.Link).
			SetURL(fd.URL).
			SetLinks(linksJson).
			SetItemDescription(feedItems).
			Save(ctx)

		if err != nil {
			fmt.Printf("failed to update: %v", err)
			return errors.New(fmt.Sprintf("failed to update: %v", err))
		}
	}

	return nil
}

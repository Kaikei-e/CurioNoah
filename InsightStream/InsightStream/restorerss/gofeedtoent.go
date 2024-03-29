package restorerss

import (
	"insightstream/domain/baseFeeds"
	"insightstream/ent"

	"github.com/mmcdole/gofeed"
)

func ExchangeToEnt(gFeeds []*gofeed.Feed) []*ent.FollowLists {
	var feedsEnt []*ent.FollowLists

	for _, feed := range gFeeds {

		var feedLinks baseFeeds.FeedLink

		var feedItems []baseFeeds.FeedItem
		for _, item := range feed.Items {
			var itemEnt baseFeeds.FeedItem
			itemEnt.ItemTitle = item.Title
			itemEnt.ItemLink = item.Link
			itemEnt.ItemDescription = item.Description
			itemEnt.Updated = item.Updated
			itemEnt.UpdatedParsed = item.UpdatedParsed
			itemEnt.Published = item.Published
			itemEnt.PublishedParsed = item.PublishedParsed
			itemEnt.GUID = item.GUID
			itemEnt.Categories = item.Categories

			feedItems = append(feedItems, itemEnt)

		}

		feedLinks.Link = feed.Links

		feedEnt := ent.FollowLists{
			Title:           feed.Title,
			Description:     feed.Description,
			Link:            feed.FeedLink,
			Language:        feed.Language,
			Links:           feedLinks,
			URL:             feed.Link,
			IsActive:        true,
			ItemDescription: feedItems,
		}

		feedsEnt = append(feedsEnt, &feedEnt)

	}

	return feedsEnt
}

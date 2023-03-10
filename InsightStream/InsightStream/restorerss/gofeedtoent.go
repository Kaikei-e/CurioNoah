package restorerss

import (
	"github.com/mmcdole/gofeed"
	"insightstream/ent"
	"insightstream/models/feeds"
)

func ExchangeToEnt(gFeeds []*gofeed.Feed) []*ent.FollowList {
	var feedsEnt []*ent.FollowList

	for _, feed := range gFeeds {

		var feedLinks feeds.FeedLink

		var feedItems []feeds.FeedItem
		for _, item := range feed.Items {
			var itemEnt feeds.FeedItem
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

		feedEnt := ent.FollowList{
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

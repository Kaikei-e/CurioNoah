package restorerss

import (
	"insightstream/domain/baseFeeds"
	"insightstream/ent"

	"github.com/mmcdole/gofeed"
)

// TODO will implement unit tests
func EntFollowListExchangeToGofeed(followLists []*ent.FollowLists) ([]*gofeed.Feed, error) {
	// To reduce the amount of data fetched per feed, set a limit on the number of items to fetch per feed
	const sendAmount = 3

	var feedList []*gofeed.Feed
	for _, feed := range followLists {

		var items []*gofeed.Item
		var authors []*gofeed.Person
		//for _, item := range feed.ItemDescription {
		for i, item := range feed.ItemDescription {
			if i >= sendAmount {
				break
			}

			if len(item.Authors) > 0 {
				for _, author := range item.Authors {
					a := &gofeed.Person{
						Name:  author,
						Email: "",
					}

					authors = append(authors, a)
				}
			}

			items = append(items, &gofeed.Item{
				Title:           item.ItemTitle,
				Link:            item.ItemLink,
				Description:     item.ItemDescription,
				Updated:         item.Updated,
				UpdatedParsed:   item.UpdatedParsed,
				Published:       item.Published,
				PublishedParsed: item.PublishedParsed,
				Authors:         authors,
				GUID:            item.GUID,
				Categories:      item.Categories,
			})
		}

		feedList = append(feedList, &gofeed.Feed{
			Title:         feed.Title,
			Description:   feed.Description,
			Link:          feed.URL,
			FeedLink:      feed.Link,
			Language:      feed.Language,
			Links:         feed.Links.Link,
			Items:         items,
			UpdatedParsed: &feed.DtUpdated,
		})
	}

	return feedList, nil
}

func ExchangeEntFeedsToGofeeds(feedsEnt []*ent.Feeds) ([]baseFeeds.EachFeed, error) {
	var convertedFeeds []baseFeeds.EachFeed

	for _, feed := range feedsEnt {
		convertedFeeds = append(convertedFeeds, baseFeeds.EachFeed{
			Id:          feed.ID,
			SiteURL:     feed.SiteURL,
			Title:       feed.Title,
			Description: feed.Description,
			FeedURL:     feed.FeedURL,
			Language:    feed.Language,
			DtCreated:   feed.DtCreated,
			DtUpdated:   feed.DtUpdated,
			Favorites:   0,
		})
	}

	return convertedFeeds, nil

}

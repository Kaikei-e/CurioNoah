package manageFeedsAmount

import (
	"github.com/mmcdole/gofeed"
)

func ReduceToLatestThreeItems(feeds []*gofeed.Feed) ([]gofeed.Feed, error) {
	const trimNum = 3
	var reducedFeeds []gofeed.Feed

	for _, feed := range feeds {
		var reducedItems []*gofeed.Item
		var reducedLinks []string

		if len(feed.Items) > trimNum {
			reducedItems = feed.Items[:trimNum]
		} else {
			reducedItems = feed.Items
		}

		if len(feed.Links) > trimNum {
			reducedLinks = feed.Links[:trimNum]
		} else {
			reducedLinks = feed.Links
		}

		feed.Items = reducedItems
		feed.Links = reducedLinks

		reducedFeeds = append(reducedFeeds, *feed)
	}

	return reducedFeeds, nil
}

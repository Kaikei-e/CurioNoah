package manageFeedsAmount

import (
	"github.com/mmcdole/gofeed"
	"sort"
)

func ReduceToLatestThreeItems(feeds []*gofeed.Feed) ([]gofeed.Feed, error) {
	const trimNum = 3
	var reducedFeeds []gofeed.Feed

	if feeds == nil {
		return nil, nil
	}

	sort.Slice(feeds, func(i, j int) bool {
		return feeds[i].UpdatedParsed.After(*feeds[j].UpdatedParsed)
	})

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

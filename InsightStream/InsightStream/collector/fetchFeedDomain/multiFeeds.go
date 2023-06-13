package fetchFeedDomain

import (
	"errors"
	"fmt"
	"insightstream/collector"

	"github.com/mmcdole/gofeed"
)

func MultiFeed(storedList []string) ([]*gofeed.Feed, error) {
	fmt.Printf("storedList: %v \n", storedList)

	// Determine initial amount of feed data to fetch per URL
	fetchFeedAmount := len(storedList)
	if fetchFeedAmount == 0 {
		return nil, errors.New("no stored URLs to fetch from")
	}

	var feeds []*gofeed.Feed

	for _, url := range storedList {
		feed, err := collector.Collector(url)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("fetch %s: %v", url, err))
		}

		// Determine the actual amount of feed items to fetch for this URL
		actualAmount := fetchFeedAmount
		if len(feed.Items) < fetchFeedAmount {
			actualAmount = len(feed.Items)
		}

		// Fetch the actual amount of feed items
		feed.Items = feed.Items[:actualAmount]

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

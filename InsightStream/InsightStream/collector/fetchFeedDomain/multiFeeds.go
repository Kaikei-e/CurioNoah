package fetchFeedDomain

import (
	"errors"
	"fmt"
	"insightstream/collector"

	"github.com/mmcdole/gofeed"
)

func MultiFeed(storedList []string) ([]*gofeed.Feed, error) {
	fmt.Printf("storedList: %v \n", storedList)

	// Determine amount of feed data to fetch per URL
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

		// If the number of items is less than the fetchFeedAmount, adjust fetchFeedAmount accordingly
		if len(feed.Items) < fetchFeedAmount {
			fetchFeedAmount = len(feed.Items)
		}
		feed.Items = feed.Items[:fetchFeedAmount]

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

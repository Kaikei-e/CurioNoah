package fetchFeedDomain

import (
	"fmt"
	"insightstream/collector"
	"time"

	"github.com/mmcdole/gofeed"
)

func MultiFeed(storedList []string) ([]*gofeed.Feed, error) {
	fmt.Printf("storedList: %v \n", storedList)

	var feeds []*gofeed.Feed

	for _, url := range storedList {
		time.Sleep(1 * time.Second)
		feed, err := collector.Collector(url)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch feeds: %w", err)
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

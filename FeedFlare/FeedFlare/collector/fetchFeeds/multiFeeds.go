package fetchFeeds

import (
	"errors"
	"feedflare/collector"
	"fmt"

	"github.com/mmcdole/gofeed"
)

func MultiFeed(storedList []string) ([]*gofeed.Feed, error) {
	var feeds []*gofeed.Feed

	for _, url := range storedList {
		feed, err := collector.Collector(url)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("fetch %s: %v", url, err))
		}

		// TODO will have functionally to check if the feed is updated or not
		feed.Items = feed.Items[:5]

		feeds = append(feeds, feed)
	}

	return feeds, nil
}

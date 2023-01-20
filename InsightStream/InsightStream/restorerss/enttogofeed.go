package restorerss

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"insightstream/ent"
)

func FeedExchange(feedsEnt []*ent.FollowList) ([]*gofeed.Feed, error) {
	// TODO will implement unit tests

	var feedList []*gofeed.Feed
	for _, feed := range feedsEnt {
		var links []string

		err := json.Unmarshal([]byte(feed.Links), &links)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("failed to unmarshal links: %v", err))
		}

		feedList = append(feedList, &gofeed.Feed{
			Title:       feed.Title,
			Description: feed.Description,
			Link:        feed.Link,
			Language:    feed.Language,
			Links:       links,
		})
	}

	return feedList, nil
}

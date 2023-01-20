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

		var items []*gofeed.Item
		var authors []*gofeed.Person
		for _, item := range feed.ItemDescription {

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
			Title:       feed.Title,
			Description: feed.Description,
			Link:        feed.Link,
			Language:    feed.Language,
			Links:       links,
			Items:       items,
		})
	}

	return feedList, nil
}

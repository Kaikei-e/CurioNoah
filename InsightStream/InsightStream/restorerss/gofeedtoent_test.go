package restorerss

import (
	"testing"

	"github.com/mmcdole/gofeed"
)

func TestExchangeToEnt(t *testing.T) {
	// Create a sample input
	gFeeds := []*gofeed.Feed{
		// Create gofeed.Feed objects with mock data
		&gofeed.Feed{
			Title:       "Feed 1",
			Description: "Description 1",
			FeedLink:    "https://example.com/feed1",
			Link:        "https://example.com",
			Language:    "en",
			Items: []*gofeed.Item{
				&gofeed.Item{
					Title:       "Item 1",
					Link:        "https://example.com/item1",
					Description: "Description of Item 1",
					Updated:     "2023-06-19T10:30:00Z",
					Published:   "2023-06-19T10:30:00Z",
					GUID:        "item1-guid",
					Categories:  []string{"category1", "category2"},
				},
			},
			// Add more gofeed.Feed objects if needed
		},
	}

	// Call the function being tested
	feedsEnt := ExchangeToEnt(gFeeds)

	// Perform assertions
	expectedFeedCount := len(gFeeds)
	if len(feedsEnt) != expectedFeedCount {
		t.Errorf("ExchangeToEnt returned %d baseFeeds, expected %d", len(feedsEnt), expectedFeedCount)
	}

	// Example assertion for the first feed
	expectedTitle := "Feed 1"
	if feedsEnt[0].Title != expectedTitle {
		t.Errorf("Feed title is %s, expected %s", feedsEnt[0].Title, expectedTitle)
	}

	// Example assertion for the first item of the first feed
	expectedItemTitle := "Item 1"
	if feedsEnt[0].ItemDescription[0].ItemTitle != expectedItemTitle {
		t.Errorf("Item title is %s, expected %s", feedsEnt[0].ItemDescription[0].ItemTitle, expectedItemTitle)
	}
}

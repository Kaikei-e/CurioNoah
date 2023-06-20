package restorerss

import (
	"insightstream/models/feeds"
	"testing"

	"insightstream/ent"
)

func TestFeedExchange(t *testing.T) {
	// Create a sample input
	feedsEnt := []*ent.FollowList{
		// Create FollowList objects with mock data
		&ent.FollowList{
			Title:       "Feed 1",
			Description: "Description 1",
			URL:         "https://example.com/feed1",
			Link:        "https://example.com",
			Language:    "en",
			ItemDescription: []feeds.FeedItem{
				// Create FeedItem objects with mock data
				feeds.FeedItem{
					ItemTitle:       "Item 1",
					ItemLink:        "https://example.com/item1",
					ItemDescription: "Item description 1",
					Updated:         "2020-01-01T00:00:00Z",
					UpdatedParsed:   nil,
					Published:       "2020-01-01T00:00:00Z",
					PublishedParsed: nil,
					Authors:         []string{"Author 1"},
					GUID:            "https://example.com/item1",
					Categories:      []string{"Category 1"},
				},
			},
			// Add more FollowList objects if needed
		},
	}

	// Call the function being tested
	feedList, err := FeedExchange(feedsEnt)

	// Perform assertions
	if err != nil {
		t.Errorf("FeedExchange returned an error: %v", err)
	}

	expectedFeedCount := len(feedsEnt)
	if len(feedList) != expectedFeedCount {
		t.Errorf("FeedExchange returned %d feeds, expected %d", len(feedList), expectedFeedCount)
	}

	// Add more assertions based on your specific requirements

	// Example assertion for the first feed
	expectedTitle := "Feed 1"
	if feedList[0].Title != expectedTitle {
		t.Errorf("Feed title is %s, expected %s", feedList[0].Title, expectedTitle)
	}

	// Example assertion for the first item of the first feed
	expectedItemTitle := "Item 1"
	if feedList[0].Items[0].Title != expectedItemTitle {
		t.Errorf("Item title is %s, expected %s", feedList[0].Items[0].Title, expectedItemTitle)
	}

}

package restorerss

import (
	"github.com/mmcdole/gofeed"
	"insightstream/models/feeds"
	"reflect"
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
	feedList, err := EntFollowListExchangeToGofeed(feedsEnt)

	// Perform assertions
	if err != nil {
		t.Errorf("EntFollowListExchangeToGofeed returned an error: %v", err)
	}

	expectedFeedCount := len(feedsEnt)
	if len(feedList) != expectedFeedCount {
		t.Errorf("EntFollowListExchangeToGofeed returned %d feeds, expected %d", len(feedList), expectedFeedCount)
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

func TestEntFollowListExchangeToGofeed(t *testing.T) {
	type args struct {
		followLists []*ent.FollowList
	}
	tests := []struct {
		name    string
		args    args
		want    []*gofeed.Feed
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EntFollowListExchangeToGofeed(tt.args.followLists)
			if (err != nil) != tt.wantErr {
				t.Errorf("EntFollowListExchangeToGofeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EntFollowListExchangeToGofeed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeEntFeedsToGofeeds(t *testing.T) {
	type args struct {
		feedsEnt []*ent.Feeds
	}
	tests := []struct {
		name    string
		args    args
		want    []feeds.EachFeed
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExchangeEntFeedsToGofeeds(tt.args.feedsEnt)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExchangeEntFeedsToGofeeds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExchangeEntFeedsToGofeeds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package manageFeedsAmount

import (
	"github.com/mmcdole/gofeed"
	"testing"
	"time"
)

func TestReduceToLatestThreeItems(t *testing.T) {
	type args struct {
		feeds []*gofeed.Feed
	}

	const trimNum = 3

	feeds := FeedGenerator()

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				feeds: feeds,
			},
			want:    trimNum,
			wantErr: false,
		},

		{
			name: "empty",
			args: args{
				feeds: nil,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "less than trimNum",
			args: args{
				feeds: []*gofeed.Feed{
					{
						Title: "test",
						Items: []*gofeed.Item{
							{
								Title: "test",
							},
						},
						Links: []string{
							"test1",
							"test2",
						},
					},
				},
			},
			want:    1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReduceToLatestThreeItems(tt.args.feeds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReduceToLatestThreeItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for i, feed := range got {
				if len(feed.Items) > trimNum {
					t.Errorf("ReduceToLatestThreeItems() feed %v has more than %v items", i, trimNum)
				}

				if len(feed.Links) > trimNum {
					t.Errorf("ReduceToLatestThreeItems() feed %v has more than %v links", i, trimNum)
				}

			}

		})
	}
}

func FeedGenerator() []*gofeed.Feed {
	const feedAmount = 5
	tNow := time.Now()

	theFeed := gofeed.Feed{
		Title: "test",
		Items: []*gofeed.Item{
			{
				Title: "test1",
			},
			{
				Title: "test2",
			},
			{
				Title: "test3",
			},
			{
				Title: "test4",
			},
			{
				Title: "test5",
			},
		},
		Links: []string{
			"site1",
			"site2",
		},
		Updated:       tNow.String(),
		UpdatedParsed: &tNow,
	}

	var feeds []*gofeed.Feed
	for i := 0; i < feedAmount; i++ {
		feeds = append(feeds, &theFeed)
	}

	return feeds
}

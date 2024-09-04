package convertToResponse

import (
	"insightstream/domain/searchedFeed"
	"insightstream/ent" // Add the import for the ent.Feeds package
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestToSearchWord(t *testing.T) {
	testFeeds := []*ent.Feeds{
		{
			Title:       "testTitle1",
			Description: "testDescription1",
			FeedURL:     "testFeedURL1",
			ID:          uuid.New(),
		},
		{
			Title:       "testTitle2",
			Description: "testDescription2",
			FeedURL:     "testFeedURL2",
			ID:          uuid.New(),
		},
	}

	type args struct {
		feeds []*ent.Feeds
	}
	tests := []struct {
		name    string
		args    args
		want    []searchedFeed.ByTitleOrDescription
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				feeds: testFeeds,
			},
			want: []searchedFeed.ByTitleOrDescription{
				{
					Title:       "testTitle1",
					Description: "testDescription1",
					FeedURL:     "testFeedURL1",
				},
				{
					Title:       "testTitle2",
					Description: "testDescription2",
					FeedURL:     "testFeedURL2",
				},
			},
			wantErr: false,
		},
		{
			name: "empty slice",
			args: args{
				feeds: nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSearchWord(tt.args.feeds)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSearchWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSearchWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

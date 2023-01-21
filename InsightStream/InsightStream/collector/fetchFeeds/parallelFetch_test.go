package fetchFeeds

import (
	"github.com/mmcdole/gofeed"
	"reflect"
	"testing"
)

func TestParallelizeFetch(t *testing.T) {
	type args struct {
		storedList []string
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
			got, err := ParallelizeFetch(tt.args.storedList)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParallelizeFetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParallelizeFetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paralyzingFetch(t *testing.T) {
	type args struct {
		feedsList []string
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
			got, err := paralyzingFetch(tt.args.feedsList)
			if (err != nil) != tt.wantErr {
				t.Errorf("paralyzingFetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paralyzingFetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

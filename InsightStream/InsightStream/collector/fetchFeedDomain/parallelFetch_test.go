package fetchFeedDomain

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelizeFetch(t *testing.T) {
	feed := `<?xml version="1.0" encoding="UTF-8"?>
	<rss version="2.0">
	<channel>
		<title>Test Feed</title>
		<description>This is a test feed</description>
		<link>http://localhost</link>
	</channel>
	</rss>`

	servers := make([]*httptest.Server, 10)
	for i := range servers {
		servers[i] = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, feed)
		}))
		defer servers[i].Close()
	}

	urls := make([]string, len(servers))
	for i, s := range servers {
		urls[i] = s.URL
	}

	feeds, err := ParallelizeFetch(urls)
	assert.NoError(t, err)

	for _, f := range feeds {
		assert.Equal(t, "Test Feed", f.Title)
		assert.Equal(t, "This is a test feed", f.Description)
		assert.Equal(t, "http://localhost", f.Link)
	}
}

// func TestParallelizeFetch(t *testing.T) {
// 	type args struct {
// 		storedList []string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []*gofeed.Feed
// 		wantErr bool
// 	}{
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := ParallelizeFetch(tt.args.storedList)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ParallelizeFetch() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ParallelizeFetch() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_paralyzingFetch(t *testing.T) {
// 	type args struct {
// 		feedsList []string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    []*gofeed.Feed
// 		wantErr bool
// 	}{
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := paralyzingFetch(tt.args.feedsList)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("paralyzingFetch() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("paralyzingFetch() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

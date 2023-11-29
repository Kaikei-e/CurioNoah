package collector

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mmcdole/gofeed"
)

func TestCollector(t *testing.T) {
	//const testXMLPath = "./testdata/sample.xml"
	const testGuardianURL = "https://www.theverge.com/rss/samsung/index.xml"

	cases := map[string]struct {
		input    string
		expected *gofeed.Feed
		wantErr  bool
	}{
		"success": {
			//input: testXMLPath,
			input: testGuardianURL,
			expected: &gofeed.Feed{
				Link: "https://www.theverge.com/samsung",
			},
			wantErr: false,
		},
		"failure": {
			input: "invalid path",
			expected: &gofeed.Feed{
				Link: "",
			},
			wantErr: true,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := Collector(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				return
			}

			if !cmp.Equal(got.Link, tt.expected.Link) {
				t.Errorf("Collector() got = %v, want %v", got, tt.expected)
			}
			//if !reflect.DeepEqual(got, tt.expected) {
			//	t.Errorf("Collector() got = %v, want %v", got.Title, tt.expected.Title)
			//}
		})
	}
}

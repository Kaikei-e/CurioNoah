package collector

import (
	"github.com/google/go-cmp/cmp"
	"github.com/mmcdole/gofeed"
	"testing"
)

func TestCollector(t *testing.T) {
	const testXMLPath = "./testdata/sample.xml"

	cases := map[string]struct {
		input    string
		expected *gofeed.Feed
		wantErr  bool
	}{
		"success": {
			input: testXMLPath,
			expected: &gofeed.Feed{
				Title: "...",
			},
			wantErr: false,
		},
		"failure": {
			input: "invalid path",
			expected: &gofeed.Feed{
				Title: "",
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

			if !cmp.Equal(got.Title, tt.expected.Title) {
				t.Errorf("Collector() got = %v, want %v", got, tt.expected)
			}
			//if !reflect.DeepEqual(got, tt.expected) {
			//	t.Errorf("Collector() got = %v, want %v", got.Title, tt.expected.Title)
			//}
		})
	}
}

package fetchFeeds

import (
	"testing"
)

func TestMultiFeed(t *testing.T) {

	const testGuardianURL = "https://rss.nytimes.com/services/xml/rss/nyt/World.xml"

	cases := map[string]struct {
		input   []string
		wantErr bool
	}{
		"success": {
			input:   []string{testGuardianURL},
			wantErr: false,
		},
		"failureWrongPath": {
			input:   []string{"invalid path"},
			wantErr: true,
		},
		"successButEmptyPath": {
			input:   []string{},
			wantErr: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := MultiFeed(c.input)
			if (err != nil) != c.wantErr {
				t.Errorf("MultiFeed() error = %v, wantErr %v", err, c.wantErr)
				return
			}
		})
	}
}

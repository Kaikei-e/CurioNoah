package registerFeed

import (
	"insightstream/ent"
	"testing"

	"github.com/mmcdole/gofeed"
)

func TestRegisterSingle(t *testing.T) {
	type args struct {
		inputLink string
		feed      *gofeed.Feed
		cl        *ent.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterSingle(tt.args.inputLink, tt.args.feed, tt.args.cl); (err != nil) != tt.wantErr {
				t.Errorf("RegisterSingle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegisterMulti(t *testing.T) {
	type args struct {
		feeds []*gofeed.Feed
		cl    *ent.Client
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterMulti(tt.args.feeds, tt.args.cl); (err != nil) != tt.wantErr {
				t.Errorf("RegisterMulti() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

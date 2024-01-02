package searchPort

import (
	"context"
	"insightstream/domain/searchWord"
	"insightstream/driver/entDriver"
	"insightstream/ent"
	"log/slog"
)

type SearchPort interface {
	SearchByTitleOrDescription(sw searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error)
}

type Impl struct{}

func (s *Impl) NewSearchPort() SearchPort {
	return &Impl{}
}

func (s *Impl) SearchByTitleOrDescription(sw searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error) {

	// change to pass a simple string
	feeds, err := entDriver.SearchFeeds(sw.Description, cl, ctx)
	if err != nil {
		return nil, err
	}
	slog.Info("baseFeeds length is", "%v", len(feeds))

	return feeds, nil
}

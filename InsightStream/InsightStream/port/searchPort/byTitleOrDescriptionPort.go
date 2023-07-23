package searchPort

import (
	"context"
	"fmt"
	"insightstream/domain/searchWord"
	"insightstream/driver/entDriver"
	"insightstream/ent"
)

type SearchPort interface {
	SearchByTitleOrDescription(sw searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error)
}

type Impl struct{}

func (s *Impl) NewSearchPort() SearchPort {
	return &Impl{}
}

func (s *Impl) SearchByTitleOrDescription(sw searchWord.SearchWord, cl *ent.Client, ctx context.Context) ([]*ent.Feeds, error) {

	feeds, err := entDriver.SearchFeeds(sw, cl, ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("baseFeeds length is : ", len(feeds))

	return feeds, nil
}

package dependencyInversion

import (
	"fmt"
	"insightstream/ent"
	"insightstream/models/feeds"
	"insightstream/repository/readfeed"
	"insightstream/restorerss"
)

type (
	FeedCollection interface {
		FetchAll(cl *ent.Client) error
		FetchSingle() error
		FetchInfinite(page int, cl *ent.Client) ([]feeds.EachFeed, bool, error)
	}

	FeedCollectionImpl struct {
	}
)

func NewFeedCollection() FeedCollection {
	return &FeedCollectionImpl{}
}

func (f *FeedCollectionImpl) FetchAll(cl *ent.Client) error {
	return nil
}

func (f *FeedCollectionImpl) FetchSingle() error {
	return nil
}

func (f *FeedCollectionImpl) FetchInfinite(page int, cl *ent.Client) ([]feeds.EachFeed, bool, error) {
	fds, hadExceeded, err := readfeed.InfiniteScroll(cl, page)
	if err != nil {
		return nil, false, fmt.Errorf("failed to query by ten: %v", err)
	}

	exchangedFeeds, err := restorerss.ExchangeEntFeedsToGofeeds(fds)
	if err != nil {
		return nil, false, err
	}

	return exchangedFeeds, hadExceeded, nil
}

package dependencyInversion

import (
	"fmt"
	"golang.org/x/exp/slices"
	"insightstream/domain/baseFeeds"
	"insightstream/driver/parser"
	"insightstream/ent"
	"insightstream/repository/readfeed"
	"insightstream/restorerss"
)

type (
	FeedCollection interface {
		FetchAll(cl *ent.Client) error
		FetchSingle() error
		FetchInfinite(page int, cl *ent.Client) ([]baseFeeds.EachFeed, bool, error)
		CompactFeeds([]baseFeeds.EachFeed) ([]baseFeeds.EachFeed, error)
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

func (f *FeedCollectionImpl) FetchInfinite(page int, cl *ent.Client) ([]baseFeeds.EachFeed, bool, error) {
	fds, hadExceeded, err := readfeed.InfiniteScroll(cl, page)
	if err != nil {
		return nil, false, fmt.Errorf("failed to query by ten: %v", err)
	}

	var formattedFeeds []*ent.Feeds
	for _, fd := range fds {
		description, err := parser.HTMLToDoc(fd.Description)
		if err != nil {
			err := fmt.Errorf("failed to parse html: %v", err)
			fmt.Println(err)
			continue
		}

		fd.Description = description
		formattedFeeds = append(formattedFeeds, fd)
	}

	exchangedFeeds, err := restorerss.ExchangeEntFeedsToGofeeds(fds)
	if err != nil {
		return nil, false, err
	}

	compactedFeeds, err := f.CompactFeeds(exchangedFeeds)
	if err != nil {
		return nil, false, err
	}

	return compactedFeeds, hadExceeded, nil
}

func (f *FeedCollectionImpl) CompactFeeds(fds []baseFeeds.EachFeed) ([]baseFeeds.EachFeed, error) {

	slices.SortStableFunc(fds, func(a, b baseFeeds.EachFeed) bool {
		return a.FeedURL < b.FeedURL
	})

	uniqueFeeds, err := removeDuplicateFeeds(fds)
	if err != nil {
		return nil, err
	}

	slices.SortStableFunc(uniqueFeeds, func(a, b baseFeeds.EachFeed) bool {
		return a.DtUpdated.After(b.DtUpdated)
	})

	return uniqueFeeds, nil
}

func removeDuplicateFeeds(fds []baseFeeds.EachFeed) ([]baseFeeds.EachFeed, error) {
	var uniqueFeeds []baseFeeds.EachFeed

	lastFeedURL := ""
	for _, feed := range fds {
		// If the current feed has a different FeedURL from the last one,
		// add it to the unique baseFeeds slice
		if feed.FeedURL != lastFeedURL {
			uniqueFeeds = append(uniqueFeeds, feed)
			lastFeedURL = feed.FeedURL
		}
	}

	return uniqueFeeds, nil
}

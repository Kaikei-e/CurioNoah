package dependencyInversion

import (
	"fmt"
	"insightstream/domain/baseFeeds"
	"insightstream/driver/parser"
	"insightstream/ent"
	"insightstream/repository/readfeed"
	"insightstream/restorerss"
	"log/slog"
	"sort"
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
			slog.Error(err.Error())
			continue
		}

		fd.Description = description
		formattedFeeds = append(formattedFeeds, fd)
	}

	exchangedFeeds, err := restorerss.ExchangeEntFeedsToGofeeds(fds)
	if err != nil {
		return nil, false, fmt.Errorf("failed to exchange ent feeds to gofeeds: %v", err)
	}

	compactedFeeds, err := f.CompactFeeds(exchangedFeeds)
	if err != nil {
		return nil, false, fmt.Errorf("failed to compact feeds: %v", err)
	}

	return compactedFeeds, hadExceeded, nil
}

func (f *FeedCollectionImpl) CompactFeeds(fds []baseFeeds.EachFeed) ([]baseFeeds.EachFeed, error) {

	sort.Slice(fds, func(i, j int) bool {
		return fds[i].FeedURL < fds[j].FeedURL
	})

	uniqueFeeds, err := removeDuplicateFeeds(fds)
	if err != nil {
		return nil, fmt.Errorf("failed to remove duplicate feeds: %w", err)
	}

	sort.Slice(uniqueFeeds, func(i, j int) bool {
		return uniqueFeeds[i].DtUpdated.After(uniqueFeeds[j].DtUpdated)
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

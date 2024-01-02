package indexing

import (
	"fmt"
	"insightstream/collector/fetchFeedDomain"
	"insightstream/ent"
	"insightstream/restorerss"
	"log/slog"
	"sync"
)

func FetchLatestByClick(cl *ent.Client) error {
	var wg sync.WaitGroup

	s := NewStoreManager(cl)

	result, err := s.FetchFollowList()
	if err != nil {
		return fmt.Errorf("failed to query all baseFeeds: %w", err)
	}

	feedExchanged, err := restorerss.EntFollowListExchangeToGofeed(result)
	if err != nil {
		return fmt.Errorf("failed to exchange feeds: %w", err)
	}

	var targetLinks []string
	for _, list := range result {
		targetLinks = append(targetLinks, list.Link)
	}

	fetchedFeeds, err := fetchFeedDomain.ParallelizeFetch(targetLinks)
	if err != nil {
		return fmt.Errorf("failed to fetch feeds: %w", err)
	}

	feedLinkList, err := CheckDiffByFeedItems(feedExchanged, fetchedFeeds)
	if err != nil {
		return fmt.Errorf("failed to check diff by feed items: %w", err)
	}

	var updateTargetList []updateList
	for _, fll := range feedLinkList {
		for _, list := range result {
			var updateTargetItem updateList
			if fll == list.Link {
				updateTargetItem = updateList{
					Link: list.Link,
					ID:   list.ID,
				}
				updateTargetList = append(updateTargetList, updateTargetItem)
				break
			}
		}
	}

	// update baseFeeds

	if len(updateTargetList) == 0 {
		return nil
	}

	followLists := restorerss.ExchangeToEnt(fetchedFeeds)
	for i, list := range followLists {
		for _, updateTarget := range updateTargetList {
			if list.Link == updateTarget.Link {
				followLists[i].ID = updateTarget.ID
			}
		}
	}

	err = s.UpdateFeeds(followLists)
	if err != nil {
		return fmt.Errorf("failed to update baseFeeds: %w", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = PingToSyncOnlyLatestFeeds()
		if err != nil {
			slog.Error("failed to ping to sync: %v", err)
			return
		}
	}()

	wg.Wait()

	return nil

}

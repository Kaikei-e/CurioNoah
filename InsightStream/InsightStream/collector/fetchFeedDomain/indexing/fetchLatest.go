package indexing

import (
	"errors"
	"insightstream/collector/fetchFeedDomain"
	"insightstream/ent"
	"insightstream/restorerss"
	"sync"

	"github.com/labstack/gommon/log"
)

func FetchLatestByClick(cl *ent.Client) error {
	var wg sync.WaitGroup

	s := NewStoreManager(cl)

	result, err := s.FetchFollowList()
	if err != nil {
		return errors.New("failed to query all baseFeeds")
	}

	feedExchanged, err := restorerss.EntFollowListExchangeToGofeed(result)
	if err != nil {
		return errors.New("failed to convert ent struct to gofeed struct")
	}

	var targetLinks []string
	for _, list := range result {
		targetLinks = append(targetLinks, list.Link)
	}

	fetchedFeeds, err := fetchFeedDomain.ParallelizeFetch(targetLinks)
	if err != nil {
		return errors.New("failed to fetch feeds")
	}

	feedLinkList, err := CheckDiffByFeedItems(feedExchanged, fetchedFeeds)
	if err != nil {
		return errors.New("failed to check diff")
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
		return errors.New("failed to update feeds")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = PingToSyncOnlyLatestFeeds()
		if err != nil {
			log.Errorf("failed to ping to sync only latest feeds: %v", err)
			return
		}
	}()

	wg.Wait()

	return nil

}

package indexing

import (
	"fmt"
	"insightstream/collector/fetchFeedDomain"
	register "insightstream/collector/registerFeed"
	"insightstream/ent"
	"insightstream/repository/readfeed"
	"insightstream/restorerss"
	"log/slog"
	"sync"

	"github.com/labstack/gommon/log"
)

type StoreManager struct {
	client *ent.Client
}

type updateList struct {
	Link string
	ID   int
}

func NewStoreManager(client *ent.Client) *StoreManager {
	return &StoreManager{client}
}

func (s *StoreManager) FetchFollowList() ([]*ent.FollowLists, error) {
	return readfeed.QueryAll(s.client)
}

func (s *StoreManager) UpdateFeeds(feeds []*ent.FollowLists) error {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := register.Update(feeds, s.client)
		if err != nil {
			slog.Error("failed to update feeds: %v", err)
		}
	}()

	wg.Wait()
	log.Info("Synchronizing baseFeeds was completed")

	return nil
}

func (s *StoreManager) Store() (*sync.WaitGroup, error) {
	var wg sync.WaitGroup

	result, err := s.FetchFollowList()
	if err != nil {
		slog.Error("failed to query all baseFeeds: %v", err)
		return nil, fmt.Errorf("failed to query all baseFeeds: %w", err)
	}

	idList, newFeeds, err := CheckDiff(result)
	if err != nil {
		return nil, fmt.Errorf("failed to check diff: %w", err)
	}

	convertedFeeds := restorerss.ExchangeToEnt(newFeeds)
	addingList := s.mergeLists(result, convertedFeeds, idList)

	if len(addingList) == 0 {
		return nil, nil
	}

	err = s.UpdateFeeds(addingList)
	if err != nil {
		slog.Error("failed to update baseFeeds: %v", err)
		return nil, fmt.Errorf("failed to update baseFeeds: %w", err)
	}

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	err = PingToSyncOnlyLatestFeeds()
	//	if err != nil {
	//		log.Errorf("failed to ping to sync only latest baseFeeds: %v", err)
	//		return
	//	}
	//}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = PingToSync()
		if err != nil {
			slog.Error("failed to ping to sync all baseFeeds: %v", err)
			return
		}
	}()

	return &wg, nil
}

func (s *StoreManager) StoreByDiff() (*sync.WaitGroup, error) {
	var wg sync.WaitGroup

	result, err := s.FetchFollowList()
	if err != nil {
		slog.Error("failed to query all baseFeeds: %v", err)
		return nil, fmt.Errorf("failed to query all baseFeeds: %w", err)
	}

	// convert existing ent struct to gofeed struct
	feedExchanged, err := restorerss.EntFollowListExchangeToGofeed(result)
	if err != nil {
		slog.Error("failed to exchange ent struct to gofeed struct: %v", err)
		return nil, fmt.Errorf("failed to exchange ent struct to gofeed struct: %w", err)
	}

	// list up target links
	var targetLinks []string
	for _, list := range result {
		targetLinks = append(targetLinks, list.Link)
	}

	fetchedFeeds, err := fetchFeedDomain.ParallelizeFetch(targetLinks)
	if err != nil {
		slog.Error("failed to fetch feeds: %v", err)
		return nil, fmt.Errorf("failed to fetch feeds: %w", err)
	}

	feedLinkList, err := CheckDiffByFeedItems(feedExchanged, fetchedFeeds)
	if err != nil {
		return nil, fmt.Errorf("failed to check diff by feed items: %w", err)
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
				fmt.Println("updateTargetItem: ", updateTargetItem)

				updateTargetList = append(updateTargetList, updateTargetItem)
				break

			}
		}

	}

	// update baseFeeds

	if len(updateTargetList) == 0 {
		return nil, nil
	}

	followLists := restorerss.ExchangeToEnt(fetchedFeeds)

	for i, list := range followLists {
		for _, updateTarget := range updateTargetList {
			if list.Link == updateTarget.Link {
				followLists[i].ID = updateTarget.ID
				break
			}
		}
	}

	err = s.UpdateFeeds(followLists)
	if err != nil {
		return nil, fmt.Errorf("failed to update baseFeeds: %w", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = PingToSyncOnlyLatestFeeds()
		if err != nil {
			slog.Error("failed to ping to sync only latest baseFeeds: %v", err)
			return
		}
	}()

	return &wg, nil

}

func (s *StoreManager) mergeLists(result []*ent.FollowLists, convertedFeeds []*ent.FollowLists, idList []int) []*ent.FollowLists {
	var addingList []*ent.FollowLists
	for _, id := range idList {
		isAdded := false
		for _, list := range result {
			if isAdded {
				break
			}
			for _, feed := range convertedFeeds {
				if id == list.ID && feed.Link == list.Link {
					// must be set because the ID is not set in the convertedFeeds
					feed.ID = list.ID
					addingList = append(addingList, feed)
					isAdded = true
					break
				}
			}
		}
	}

	return addingList
}

//
//func Store(cl *ent.Client) error {
//	var wg sync.WaitGroup
//
//	result, err := readfeed.QueryAll(cl)
//	if err != nil {
//		return errors.New(fmt.Sprintf("failed to query all: %v", err))
//	}
//
//	fmt.Printf("result: %v \n", len(result))
//
//	idList, newFeeds, err := CheckDiff(result)
//	if err != nil {
//		return errors.New(fmt.Sprintf("failed to check diff: %v", err))
//	}
//
//	sort.SliceStable(result, func(i, j int) bool {
//		return result[i].Link < result[j].Link
//	})
//
//	sort.SliceStable(newFeeds, func(i, j int) bool {
//		return newFeeds[i].Link < newFeeds[j].Link
//	})
//
//	// convert gofeed.Feed to ent.FollowLists
//	convertedFeeds := restorerss.ExchangeToEnt(newFeeds)
//
//	var addingList []*ent.FollowLists
//	for _, id := range idList {
//
//		isAdded := false
//
//		for _, list := range result {
//			if isAdded {
//				break
//			}
//
//			for _, feed := range convertedFeeds {
//				if id == list.ID && feed.Link == list.Link {
//
//					// must be set because the ID is not set in the convertedFeeds
//					feed.ID = list.ID
//
//					addingList = append(addingList, feed)
//					isAdded = true
//					break
//				}
//			}
//		}
//	}
//
//	if len(addingList) == 0 {
//		return nil
//	}
//
//	wg.Add(1)
//
//	go func() {
//		defer wg.Done()
//		err = register.Update(addingList, cl)
//		if err != nil {
//			log.Warnj(map[string]interface{}{
//				"error": err,
//			})
//
//		}
//	}()
//
//	wg.Wait()
//	log.Info("Synchronizing baseFeeds was completed")
//
//	err = PingToSync()
//	if err != nil {
//		log.Errorf("failed to ping to sync: %v", err)
//		return err
//	}
//
//	return nil
//}

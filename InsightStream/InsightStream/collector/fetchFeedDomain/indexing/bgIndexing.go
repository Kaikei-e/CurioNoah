package indexing

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"insightstream/collector/fetchFeedDomain"
	register "insightstream/collector/registerFeed"
	"insightstream/ent"
	"insightstream/repository/readfeed"
	"insightstream/restorerss"
	"sync"
)

type StoreManager struct {
	client *ent.Client
}

func NewStoreManager(client *ent.Client) *StoreManager {
	return &StoreManager{client}
}

func (s *StoreManager) FetchFollowListss() ([]*ent.FollowLists, error) {
	return readfeed.QueryAll(s.client)
}

func (s *StoreManager) UpdateFeeds(feeds []*ent.FollowLists) error {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := register.Update(feeds, s.client)
		if err != nil {
			log.Warnj(map[string]interface{}{
				"error": err,
			})
		}
	}()

	wg.Wait()
	log.Info("Synchronizing feeds was completed")

	return nil
}

func (s *StoreManager) Store() (*sync.WaitGroup, error) {
	var wg sync.WaitGroup

	result, err := s.FetchFollowListss()
	if err != nil {
		return nil, errors.New("failed to query all feeds")
	}

	idList, newFeeds, err := CheckDiff(result)
	if err != nil {
		return nil, errors.New("failed to check diff")
	}

	convertedFeeds := restorerss.ExchangeToEnt(newFeeds)
	addingList := s.mergeLists(result, convertedFeeds, idList)

	if len(addingList) == 0 {
		return nil, nil
	}

	err = s.UpdateFeeds(addingList)
	if err != nil {
		log.Errorf("failed to update feeds: %v", err)
		return nil, errors.New("failed to update feeds")
	}

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	err = PingToSyncOnlyLatestFeeds()
	//	if err != nil {
	//		log.Errorf("failed to ping to sync only latest feeds: %v", err)
	//		return
	//	}
	//}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = PingToSync()
		if err != nil {
			log.Errorf("failed to ping to sync all feeds: %v", err)
			return
		}
	}()

	return &wg, nil
}

func (s *StoreManager) StoreByDiff() (*sync.WaitGroup, error) {
	var wg sync.WaitGroup

	result, err := s.FetchFollowListss()
	if err != nil {
		return nil, errors.New("failed to query all feeds")
	}

	// convert existing ent struct to gofeed struct
	feedExchanged, err := restorerss.EntFollowListExchangeToGofeed(result)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to excahnge ent to gofeed struct. error: %v", err))
	}

	// list up target links
	var targetLinks []string
	for _, list := range result {
		targetLinks = append(targetLinks, list.Link)
	}

	fetchedFeeds, err := fetchFeedDomain.ParallelizeFetch(targetLinks)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to fetch feed. error: %v", err))
	}

	feedLinkList, err := CheckDiffByFeedItems(feedExchanged, fetchedFeeds)

	var updateTargetIDList []int
	for _, fll := range feedLinkList {
		for _, list := range result {
			if fll == list.Link {
				updateTargetIDList = append(updateTargetIDList, list.ID)
			}
		}

	}

	// update feeds
	fmt.Println("updateTargetIDList: ", updateTargetIDList)

	if len(updateTargetIDList) == 0 {
		return nil, nil
	}
	FollowLists := restorerss.ExchangeToEnt(fetchedFeeds)

	err = s.UpdateFeeds(FollowLists)
	if err != nil {
		log.Errorf("failed to update feeds: %v", err)
		return nil, errors.New("failed to update feeds")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = PingToSync()
		if err != nil {
			log.Errorf("failed to ping to sync all feeds: %v", err)
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
//	log.Info("Synchronizing feeds was completed")
//
//	err = PingToSync()
//	if err != nil {
//		log.Errorf("failed to ping to sync: %v", err)
//		return err
//	}
//
//	return nil
//}

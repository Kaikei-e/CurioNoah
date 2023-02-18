package indexing

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	register "insightstream/collector/registerFeed"
	"insightstream/ent"
	"insightstream/repository/readfeed"
	"insightstream/restorerss"
	"sort"
)

//const (
//	routineInterval = 60 * time.Second
//)

func Store(cl *ent.Client) error {
	result, err := readfeed.QueryAll(cl)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to query all: %v", err))
	}

	fmt.Printf("result: %v \n", len(result))

	idList, newFeeds, err := CheckDiff(result)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to check diff: %v", err))
	}

	// Comment out this snippet:
	// because I can't remember why I wrote this snippet.
	//for _, list := range result {
	//	n := time.Now()
	//
	//	if list.DtLastInserted.Before(n.Add(-(routineInterval + 45*time.Second))) {
	//		fmt.Sprintf("list: %v", list)
	//		links = append(links, list.Link)
	//		//fmt.Println("already updated")
	//		//return
	//	}
	//}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Link < result[j].Link
	})

	sort.SliceStable(newFeeds, func(i, j int) bool {
		return newFeeds[i].Link < newFeeds[j].Link
	})

	// convert gofeed.Feed to ent.FollowList
	convertedFeeds := restorerss.ExchangeToEnt(newFeeds)

	var addingList []*ent.FollowList
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

	if len(addingList) == 0 {
		return nil
	}

	err = register.Update(addingList, cl)
	if err != nil {
		log.Warnj(map[string]interface{}{
			"error": err,
		})

		return errors.New(fmt.Sprintf("failed to update the RSS feed list"))
	}

	return nil
}

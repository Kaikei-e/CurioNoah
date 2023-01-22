package realtime

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/mmcdole/gofeed"
	"insightstream/collector/fetchFeeds"
	"insightstream/ent"
	"insightstream/restorerss"
	"sort"
)

// maxLinks is the maximum number of links to be fetched
// and stored in the database
const maxLinks = 3

func CheckDiff(fl []*ent.FollowList) ([]int, error) {
	fmt.Printf("fl: %v \n", len(fl))

	var links []string
	for _, list := range fl {
		links = append(links, list.Link)
	}

	oldFeeds, err := restorerss.FeedExchange(fl)
	if err != nil {
		// TODO
		fmt.Errorf("failed to excahnge ent to gofeed struct. error: %v", err)
		return nil, errors.New(fmt.Sprintf("failed to excahnge ent to gofeed struct. error: %v", err))
	}

	var newFeeds []*gofeed.Feed
	feed, err := fetchFeeds.ParallelizeFetch(links)
	if err != nil {
		// TODO
		fmt.Errorf("failed to fetch feeds. error: %v", err)
		return nil, errors.New(fmt.Sprintf("failed to fetch feed. error: %v", err))
	}

	newFeeds = append(newFeeds, feed...)

	//newFeeds, err := fetchFeeds.MultiFeed(fl)
	//if err != nil {
	//	// TODO
	//	fmt.Errorf("failed to fetch feeds. error: %v", err)
	//	return nil, errors.New(fmt.Sprintf("failed to fetch feed. error: %v", err))
	//}

	sort.SliceStable(oldFeeds, func(i, j int) bool {
		return oldFeeds[i].Link < oldFeeds[j].Link
	})

	sort.SliceStable(newFeeds, func(i, j int) bool {
		return newFeeds[i].Link < newFeeds[j].Link
	})

	// TODO need to shallow the nested loop
	// This loop is too steep
	var updateList []int
	for i, feed := range fl {
		if !cmp.Equal(feed.Links.Link, newFeeds[i].Links) {
			updateList = append(updateList, feed.ID)
		}
	}

	fmt.Printf("updateList: %v \n", len(updateList))

	return updateList, nil

}

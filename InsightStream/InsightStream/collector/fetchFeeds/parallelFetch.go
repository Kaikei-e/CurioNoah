package fetchFeeds

import (
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"sync"
)

func ParallelizeFetch(storedList []string) ([]*gofeed.Feed, error) {
	var separatedList [][]string

	parallelAmount := len(storedList) / 5
	div0 := len(storedList) % 5

	if div0 != 0 {
		parallelAmount++
	}

	for i := 0; i < parallelAmount; i++ {
		if i == parallelAmount-1 {
			separatedList = append(separatedList, storedList[i*5:])
		} else {
			separatedList = append(separatedList, storedList[i*5:(i+1)*5])
		}
	}

	var parallelList [][]*gofeed.Feed

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for _, list := range separatedList {
			paralyzedFeeds, err := paralyzingFetch(list)
			if err != nil {
				// TODO have to consider error handling in this go routine
				fmt.Println(err)
				errors.New(fmt.Sprintf("failed to parallelize list: %v", err))
			}

			parallelList = append(parallelList, paralyzedFeeds)
		}
	}()

	wg.Wait()

	var flattenedList []*gofeed.Feed
	for _, feeds := range parallelList {
		for _, feed := range feeds {
			flattenedList = append(flattenedList, feed)
		}
	}

	return flattenedList, nil
}

func paralyzingFetch(feedsList []string) ([]*gofeed.Feed, error) {

	feeds, err := MultiFeed(feedsList)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to fetch feeds: %v", err))
	}

	return feeds, err
}

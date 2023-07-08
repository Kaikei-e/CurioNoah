package fetchFeedDomain

import (
	"fmt"
	"sync"

	"github.com/mmcdole/gofeed"
)

func ParallelizeFetch(storedList []string) ([]*gofeed.Feed, error) {
	// one of the most complicated part of the code
	// this part is to make sure that the feeds are fetched in parallel
	var wg sync.WaitGroup
	feedsCh := make(chan []*gofeed.Feed, len(storedList))
	errCh := make(chan error, len(storedList))

	const partition = 5

	m := &sync.Mutex{} // Mutex to protect critical section

	for i := 0; i < len(storedList); i += partition {
		end := i + partition
		if end > len(storedList) {
			end = len(storedList)
		}

		wg.Add(1)
		go func(list []string) {
			defer wg.Done()
			feeds, err := MultiFeed(list)
			if err != nil {
				errCh <- fmt.Errorf("failed to fetch feeds: %w", err)
				return
			}
			m.Lock() // Locking
			feedsCh <- feeds
			m.Unlock() // Unlocking
		}(storedList[i:end])
	}

	wg.Wait()
	close(feedsCh)
	close(errCh)

	if len(errCh) > 0 {
		return nil, <-errCh
	}

	var feeds []*gofeed.Feed
	for f := range feedsCh {
		feeds = append(feeds, f...)
	}

	return feeds, nil
}

// func ParallelizeFetch(storedList []string) ([]*gofeed.Feed, error) {
// 	var separatedList [][]string

// 	parallelAmount := len(storedList) / 5
// 	div0 := len(storedList) % 5

// 	if div0 != 0 {
// 		parallelAmount++
// 	}

// 	for i := 0; i < parallelAmount; i++ {
// 		if i == parallelAmount-1 {
// 			separatedList = append(separatedList, storedList[i*5:])
// 		} else {
// 			separatedList = append(separatedList, storedList[i*5:(i+1)*5])
// 		}
// 	}

// 	var parallelList [][]*gofeed.Feed

// 	// one of the most complicated part of the code
// 	// this part is to make sure that the feeds are fetched in parallel
// 	var wg sync.WaitGroup
// 	var ch = make(chan [][]*gofeed.Feed, len(separatedList))
// 	for _, list := range separatedList {
// 		wg.Add(1)
// 		go func(list []string) {
// 			defer wg.Done()
// 			feeds, err := paralyzingFetch(list)
// 			if err != nil {
// 				panic(err)
// 			}
// 			ch <- [][]*gofeed.Feed{feeds}
// 		}(list)

// 		parallelList = append(parallelList, <-ch...)

// 	}

// 	fmt.Println("len is ", len(parallelList))

// 	wg.Wait()

// 	var flattenedList []*gofeed.Feed
// 	for _, feeds := range parallelList {
// 		flattenedList = append(flattenedList, feeds...)
// 	}

// 	return flattenedList, nil
// }

// func paralyzingFetch(feedsList []string) ([]*gofeed.Feed, error) {
// 	//li := <-feedsList
// 	feeds, err := MultiFeed(feedsList)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("failed to fetch feeds: %v", err))
// 	}

// 	return feeds, err
// }

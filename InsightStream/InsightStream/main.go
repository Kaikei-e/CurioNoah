package main

import (
	"fmt"
	"insightstream/collector/fetchFeedDomain/indexing"
	"insightstream/domain/groupDataPool"
	"insightstream/repository"
	"insightstream/server"
	"sync"
	"time"
)

func main() {
	// 20 minute is not intentional value. Just for testing.
	ticker := time.NewTicker(20 * time.Minute)
	done := make(chan bool)

	var wg sync.WaitGroup

	cl := repository.InitConnection()
	go func() {
		for {
			wg.Add(1)
			select {
			case <-done:
				return
			case <-ticker.C:
				err := indexing.Store(cl)
				if err != nil {
					// TODO wil add logger
					err := fmt.Sprintf("failed to store: %v", err)
					fmt.Println(err)
				}

				wg.Done()
				wg.Wait()

				err = groupDataPool.UpdateGroupDataPool()
				if err != nil {
					// TODO wil add logger
					err := fmt.Sprintf("failed to update group data pool: %v", err)
					fmt.Println(err)
				}

			}
		}
	}()

	server.Server(cl)

}

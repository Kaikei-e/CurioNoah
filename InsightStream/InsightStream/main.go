package main

import (
	"fmt"
	"insightstream/collector/fetchFeedDomain/indexing"
	"insightstream/domain/groupDataPool"
	"insightstream/repository"
	"insightstream/server"
	"time"
)

func main() {
	// 20 minute is not intentional value. Just for testing.
	ticker := time.NewTicker(20 * time.Minute)
	done := make(chan bool)

	cl := repository.InitConnection()

	storeManager := indexing.NewStoreManager(cl)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				wg, err := storeManager.Store()
				if err != nil {
					// TODO will add logger
					fmt.Println(fmt.Sprintf("failed to store: %v", err))
				} else if wg != nil {
					wg.Wait() // wait for Store() to finish

					err = groupDataPool.UpdateGroupDataPool()
					if err != nil {
						// TODO will add logger
						fmt.Println(fmt.Sprintf("failed to update group data pool: %v", err))
					}
				}
			}
		}
	}()

	server.Server(cl)

}

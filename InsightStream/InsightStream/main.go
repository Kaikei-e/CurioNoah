//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"insightstream/collector/fetchFeedDomain/indexing"
	"insightstream/domain/groupDataPool"
	"insightstream/repository"
	"insightstream/server"
	"log"
	"log/slog"
	"os"
	"time"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))
}

func main() {
	// 4 hours is not intentional value. Just for saving resources.
	ticker := time.NewTicker(4 * time.Hour)
	done := make(chan bool)

	cl := repository.InitConnection()

	storeManager := indexing.NewStoreManager(cl)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				//wg, err := storeManager.Store()
				wg, err := storeManager.StoreByDiff()
				if err != nil {
					// TODO will add logger
					fmt.Println(fmt.Sprintf("failed to store: %v", err))
				} else if wg != nil {
					wg.Wait() // wait for Store() to finish

					log.Println("store finished")

					err = groupDataPool.UpdateGroupDataPool()
					if err != nil {
						// TODO will add logger
						fmt.Println(fmt.Sprintf("failed to update group data pool: %v", err))
					}

					log.Println("update group data pool finished")
				}
			}
		}
	}()

	server.Server(cl)

}

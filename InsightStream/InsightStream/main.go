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
	"time"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	// 360 minute is not intentional value. Just for saving resources.
	ticker := time.NewTicker(360 * time.Minute)
	done := make(chan bool)

	cl := repository.InitConnection()

	storeManager := indexing.NewStoreManager(cl)

	err := entc.Generate("./ent/schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

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

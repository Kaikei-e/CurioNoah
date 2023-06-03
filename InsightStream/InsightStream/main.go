package main

import (
	"fmt"
	"insightstream/collector/fetchFeedDomain/indexing"
	"insightstream/repository"
	"insightstream/server"
	"time"
)

func main() {
	// 30 minute is not intentional value. Just for testing.
	ticker := time.NewTicker(30 * time.Minute)
	done := make(chan bool)

	cl := repository.InitConnection()
	go func() {
		for {
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
			}
		}
	}()

	server.Server(cl)

}

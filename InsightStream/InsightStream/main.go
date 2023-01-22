package main

import (
	"insightstream/collector/fetchFeeds/realtime"
	"insightstream/repository"
	"insightstream/server"
	"time"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)
	done := make(chan bool)

	cl := repository.InitConnection()
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				realtime.Store(cl)
			}
		}
	}()

	server.Server(cl)

}

package main

import (
	"feedflare/collector"
	"feedflare/server"
	"fmt"
)

func main() {
	//const targetFilePath = "./models/feeds/sample.xml"
	const testGuardianURL = "https://rss.nytimes.com/services/xml/rss/nyt/World.xml"

	feed, err := collector.Collector(testGuardianURL)
	if err != nil {
		// TODO fix error handling
		panic(err)

	}

	fmt.Println(feed)

	server.Server()

}

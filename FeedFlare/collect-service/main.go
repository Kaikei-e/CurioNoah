package main

import (
	"feedflare/collector"
	"fmt"
)

func main() {
	const targetFilePath = "./models/feeds/sample.xml"

	feed, err := collector.Collector(targetFilePath)
	if err != nil {
		// TODO fix error handling
		panic(err)

	}

	fmt.Println(feed)
}

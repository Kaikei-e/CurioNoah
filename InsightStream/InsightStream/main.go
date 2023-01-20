package main

import (
	"insightstream/server"
)

func main() {
	//const targetFilePath = "./models/feeds/sample.xml"
	//const testGuardianURL = "https://rss.nytimes.com/services/xml/rss/nyt/World.xml"
	//
	//_, err := collector.Collector(testGuardianURL)
	//if err != nil {
	//	// TODO fix error handling
	//	panic(err)
	//
	//}

	server.Server()

}

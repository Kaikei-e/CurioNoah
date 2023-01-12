package collector

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"io"
	"os"
)

func Collector() {
	const targetFilePath = "./models/feeds/sample.xml"

	f, err := os.OpenFile(targetFilePath, os.O_RDWR, 0644)
	if err != nil {
		//TODO fix error handling
		panic(err)
	}

	redByte, err := io.ReadAll(f)
	if err != nil {
		//TODO fix error handling
		panic(err)
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(redByte))
	//feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)
	fmt.Println(feed.Items[0].Title)

	for _, author := range feed.Authors {
		fmt.Println(author.Name)
	}
}

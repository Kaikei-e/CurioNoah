package collector

import (
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"io"
	"os"
)

func Collector(targetFilePath string) (*gofeed.Feed, error) {

	f, err := os.OpenFile(targetFilePath, os.O_RDWR, 0644)
	if err != nil {
		//TODO fix error handling
		return nil, errors.New(fmt.Sprintf("open %s: %v", targetFilePath, err))
	}

	redByte, err := io.ReadAll(f)
	if err != nil {
		//TODO fix error handling
		return nil, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(redByte))
	//feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println(feed.Title)
	fmt.Println(feed.Items[0].Title)

	for _, author := range feed.Authors {
		fmt.Println(author.Name)
	}

	return feed, nil
}

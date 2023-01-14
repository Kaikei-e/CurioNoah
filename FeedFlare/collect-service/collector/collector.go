package collector

import (
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
)

func Collector(targetURL string) (*gofeed.Feed, error) {

	//f, err := os.OpenFile(targetURL, os.O_RDWR, 0644)
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("open %s: %v", targetURL, err))
	//}
	//
	//redByte, err := io.ReadAll(f)
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("read %s: %v", targetURL, err))
	//}

	fp := gofeed.NewParser()
	//feed, err := fp.ParseString(string(redByte))
	////if err != nil {
	////	return nil, errors.New(fmt.Sprintf("parse %s: %v", targetURL, err))
	////}
	feed, err := fp.ParseURL(targetURL)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("parse %s: %v", targetURL, err))
	}

	for _, author := range feed.Authors {
		fmt.Println(author.Name)
	}

	return feed, nil
}

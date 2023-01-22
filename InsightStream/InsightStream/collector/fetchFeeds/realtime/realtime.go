package realtime

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"insightstream/collector/fetchFeeds"
	register "insightstream/collector/registerFeed"
	"insightstream/ent"
	"insightstream/repository/readfeed"
	"time"
)

const (
	routineInterval = 15 * time.Second
)

func Store(cl *ent.Client) {
	result, err := readfeed.QueryAll(cl)
	if err != nil {
		errors.New(fmt.Sprintf("failed to query all: %v", err))
	}

	fmt.Printf("result: %v \n", len(result))

	var links []string
	for _, list := range result {
		n := time.Now()

		if list.DtLastInserted.Before(n.Add(-(routineInterval + 45*time.Second))) {
			fmt.Sprintf("list: %v", list)
			links = append(links, list.Link)
			//fmt.Println("already updated")
			//return
		}
	}
	idList, err := CheckDiff(result)
	if err != nil {
		errors.New(fmt.Sprintf("failed to check diff: %v", err))
	}

	fmt.Println("//////////////////////////////////")

	for i, _ := range idList {
		links = append(links, result[i].Link)
	}

	log.Infoj(map[string]interface{}{
		"links":        links,
		"all fetching": "start",
	})

	feeds, err := fetchFeeds.ParallelizeFetch(links)
	if err != nil {
		log.Warnj(map[string]interface{}{
			"error": err,
		})
	}

	log.Infoj(map[string]interface{}{
		"feed length is ": len(feeds),
		"all fetching":    "end",
	})

	if len(feeds) == 0 {
		return
	}

	err = register.Update(feeds, cl)
	if err != nil {
		log.Warnj(map[string]interface{}{
			"error": err,
		})

	}
}

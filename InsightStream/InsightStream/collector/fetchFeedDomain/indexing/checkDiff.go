package indexing

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/mmcdole/gofeed"
	"golang.org/x/exp/slices"
	"insightstream/collector/fetchFeedDomain"
	"insightstream/ent"
	"insightstream/restorerss"
	"sort"
)

// maxLinks is the maximum number of links to be fetched
// and stored in the database
//const maxLinks = 3

type linksItem struct {
	Links []string
	URL   string
}

func CheckDiff(fl []*ent.FollowList) ([]int, []*gofeed.Feed, error) {
	fmt.Printf("fl: %v \n", len(fl))

	var links []string
	for _, list := range fl {
		links = append(links, list.Link)
	}

	oldFeeds, err := restorerss.FeedExchange(fl)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("failed to excahnge ent to gofeed struct. error: %v", err))
	}

	newFeeds, err := fetchFeedDomain.ParallelizeFetch(links)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("failed to fetch feed. error: %v", err))
	}

	// create link group of old feeds that is stored in the database
	var oldLinks []linksItem
	for _, feed := range oldFeeds {
		var sortedLink []string
		var oldLink linksItem

		for _, link := range feed.Links {
			sortedLink = append(sortedLink, link)
		}

		sort.SliceStable(sortedLink, func(i, j int) bool {
			return sortedLink[i] < sortedLink[j]
		})

		oldLink.Links = append(oldLink.Links, sortedLink...)
		oldLink.URL = feed.Link

		oldLinks = append(oldLinks, oldLink)
	}

	// caution: newFeeds[x].Links is not target URLs. So need to use .Items.Link
	// create link group of new feeds that is fetched from the internet
	var newLinks []linksItem
	for _, feed := range newFeeds {
		var sortedLink []string
		var newLink linksItem

		for _, link := range feed.Items {
			sortedLink = append(sortedLink, link.Link)
		}

		sort.SliceStable(sortedLink, func(i, j int) bool {
			return sortedLink[i] < sortedLink[j]
		})

		newLink.Links = append(newLink.Links, sortedLink...)
		newLink.URL = feed.Link

		newLinks = append(newLinks, newLink)
	}

	for _, oLink := range oldLinks {
		sort.SliceStable(oLink.Links, func(i, j int) bool {
			return oLink.Links[i] < oLink.Links[j]
		})
	}

	for _, nLink := range newLinks {
		sort.SliceStable(nLink.Links, func(i, j int) bool {
			return nLink.Links[i] < nLink.Links[j]
		})
	}

	// compare oldLinks and newLinks
	var updateLinkList []string
	for i, link := range oldLinks {
		for i2, l := range link.Links {
			if cmp.Diff(l, newLinks[i].Links[i2]) != "" {
				updateLinkList = append(updateLinkList, fl[i].Link)
				break
			}
		}
	}

	// to compare oldLinks and newLinks
	// and return the list of links that need to be updated
	//var updateLinkList []string
	//for _, oLink := range oldLinks {
	//	for _, nl := range newLinks {
	//		if oLink.URL == nl.URL {
	//			for _, link := range oLink.Links {
	//				for i, nLink := range nl.Links {
	//					if link != nLink {
	//						updateLinkList = append(updateLinkList, fl[i].Link)
	//						break
	//					}
	//				}
	//			}
	//		}
	//	}
	//}

	var updateIDList []int
	for _, list := range fl {
		for _, ul := range updateLinkList {
			if list.Link == ul {
				updateIDList = append(updateIDList, list.ID)
			}
		}
	}

	compactedIDList := slices.Compact(updateIDList)

	fmt.Printf("updateIDList: %v \n", compactedIDList)

	return compactedIDList, newFeeds, nil

}

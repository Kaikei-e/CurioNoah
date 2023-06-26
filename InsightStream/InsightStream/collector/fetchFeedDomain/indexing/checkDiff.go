package indexing

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/mmcdole/gofeed"
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
	ID    int
}

func CheckDiff(fl []*ent.FollowList) ([]int, []*gofeed.Feed, error) {

	// convert existing ent struct to gofeed struct
	feedExchanged, err := restorerss.EntFollowListExchangeToGofeed(fl)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("failed to excahnge ent to gofeed struct. error: %v", err))
	}

	// list up target links
	var targetLinks []string
	for _, list := range fl {
		targetLinks = append(targetLinks, list.Link)
	}

	fetchedFeeds, err := fetchFeedDomain.ParallelizeFetch(targetLinks)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("failed to fetch feed. error: %v", err))
	}

	var fetchedLinks []linksItem
	for _, feed := range fetchedFeeds {
		var sortedLink []string
		var fetchedLink linksItem

		for _, link := range feed.Links {
			sortedLink = append(sortedLink, link)
		}

		sort.SliceStable(sortedLink, func(i, j int) bool {
			return sortedLink[i] < sortedLink[j]
		})

		// you must identify the difference between the feed.Link and the feed.FeedLink
		// feed.FeedLink is the URL of the RSS feed itself
		fetchedLink.Links = append(fetchedLink.Links, sortedLink...)
		fetchedLink.URL = feed.Link

		fetchedLinks = append(fetchedLinks, fetchedLink)
	}

	var exchangedLinks []linksItem
	for i, feed := range feedExchanged {
		var sortedLink []string
		var fetchedLink linksItem

		for _, link := range feed.Links {
			sortedLink = append(sortedLink, link)
		}

		sort.SliceStable(sortedLink, func(i, j int) bool {
			return sortedLink[i] < sortedLink[j]
		})

		fetchedLink.Links = append(fetchedLink.Links, sortedLink...)
		fetchedLink.URL = fetchedFeeds[i].Link

		exchangedLinks = append(exchangedLinks, fetchedLink)

	}

	var updateLinkList []int
	for _, fetchedLink := range fetchedLinks {
		// We look for the corresponding item in exchangedLinks
		for _, exchangedLink := range exchangedLinks {
			found := false
			if !cmp.Equal(fetchedLink.Links, exchangedLink.Links) {
				for i, list := range fl {
					if list.URL == fetchedLink.URL {
						updateLinkList = append(updateLinkList, fl[i].ID)
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}
	}

	sort.Ints(updateLinkList)
	fmt.Println("updateLinkList: ", updateLinkList)

	return updateLinkList, fetchedFeeds, nil

}

//func CheckDiff(fl []*ent.FollowList) ([]int, []*gofeed.Feed, error) {
//	fmt.Printf("fl: %v \n", len(fl))
//
//	var links []string
//	for _, list := range fl {
//		links = append(links, list.Link)
//	}
//
//	oldFeeds, err := restorerss.EntFollowListExchangeToGofeed(fl)
//	if err != nil {
//		return nil, nil, errors.New(fmt.Sprintf("failed to excahnge ent to gofeed struct. error: %v", err))
//	}
//
//	newFeeds, err := fetchFeedDomain.ParallelizeFetch(links)
//	if err != nil {
//		return nil, nil, errors.New(fmt.Sprintf("failed to fetch feed. error: %v", err))
//	}
//
//	// create link group of old feeds that is stored in the database
//	var oldLinks []linksItem
//	for _, feed := range oldFeeds {
//		var sortedLink []string
//		var oldLink linksItem
//
//		sortedLink = append(sortedLink, feed.Links...)
//
//		sort.SliceStable(sortedLink, func(i, j int) bool {
//			return sortedLink[i] < sortedLink[j]
//		})
//
//		oldLink.Links = append(oldLink.Links, sortedLink...)
//		oldLink.URL = feed.Link
//
//		oldLinks = append(oldLinks, oldLink)
//	}
//
//	// caution: newFeeds[x].Links is not target URLs. So need to use .Items.Link
//	// create link group of new feeds that is fetched from the internet
//	var newLinks []linksItem
//	for _, feed := range newFeeds {
//		var sortedLink []string
//		var newLink linksItem
//
//		for _, link := range feed.Items {
//			sortedLink = append(sortedLink, link.Link)
//		}
//
//		sort.SliceStable(sortedLink, func(i, j int) bool {
//			return sortedLink[i] < sortedLink[j]
//		})
//
//		newLink.Links = append(newLink.Links, sortedLink...)
//		newLink.URL = feed.Link
//
//		newLinks = append(newLinks, newLink)
//	}
//
//	for _, oLink := range oldLinks {
//		sort.SliceStable(oLink.Links, func(i, j int) bool {
//			return oLink.Links[i] < oLink.Links[j]
//		})
//	}
//
//	for _, nLink := range newLinks {
//		sort.SliceStable(nLink.Links, func(i, j int) bool {
//			return nLink.Links[i] < nLink.Links[j]
//		})
//	}
//
//	//var sortedOldLinks []linksItem
//	//for _, oLink := range oldLinks {
//	//	sort.Sort(sort.StringSlice(oLink.Links))
//	//	sortedOldLinks = append(sortedOldLinks, oLink)
//	//}
//	//
//	//var sortedNewLinks []linksItem
//	//for _, nLink := range newLinks {
//	//	sort.Sort(sort.StringSlice(nLink.Links))
//	//	sortedNewLinks = append(sortedNewLinks, nLink)
//	//}
//
//	// compare oldLinks and newLinks
//	var updateLinkList []string
//	for i, link := range oldLinks {
//		for i2, l := range link.Links {
//			if i2 < len(newLinks[i].Links) && cmp.Diff(l, newLinks[i].Links[i2]) != "" {
//				updateLinkList = append(updateLinkList, fl[i].Link)
//				break
//			}
//		}
//	}
//
//	var updateIDList []int
//	for _, list := range fl {
//		for _, ul := range updateLinkList {
//			if list.Link == ul {
//				updateIDList = append(updateIDList, list.ID)
//			}
//		}
//	}
//
//	compactedIDList := slices.Compact(updateIDList)
//
//	fmt.Printf("updateIDList: %v \n", compactedIDList)
//
//	return compactedIDList, newFeeds, nil
//
//}

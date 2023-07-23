package convertToResponse

import (
	"insightstream/domain/searchedFeed"
	"insightstream/ent"
)

func ToSearchWord(feeds []*ent.Feeds) ([]searchedFeed.ByTitleOrDescription, error) {
	var sfByTitleOrDescription []searchedFeed.ByTitleOrDescription

	for _, feed := range feeds {
		var sf searchedFeed.ByTitleOrDescription

		sf.Title = feed.Title
		sf.Description = feed.Description
		sf.FeedURL = feed.FeedURL

		sfByTitleOrDescription = append(sfByTitleOrDescription, sf)
	}

	return sfByTitleOrDescription, nil
}

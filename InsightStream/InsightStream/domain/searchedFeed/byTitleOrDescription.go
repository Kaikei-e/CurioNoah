package searchedFeed

type ByTitleOrDescription struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FeedURL     string `json:"feed_url"`
}

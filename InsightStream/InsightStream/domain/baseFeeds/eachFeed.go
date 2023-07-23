package baseFeeds

import (
	"github.com/google/uuid"
	"time"
)

type EachFeed struct {
	Id          uuid.UUID `json:"item_id"`
	SiteURL     string    `json:"site_url,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FeedURL     string    `json:"feed_url"`
	Language    string    `json:"language,omitempty"`
	DtCreated   time.Time `json:"dt_created"`
	DtUpdated   time.Time `json:"dt_updated"`
	Favorites   int       `json:"favorites"`
}

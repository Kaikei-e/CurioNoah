package feeds

import (
	"github.com/google/uuid"
	"time"
)

type FeedItem struct {
	Id              uuid.UUID  `json:"item_id"`
	ItemDescription string     `json:"item_description"`
	ItemLink        string     `json:"item_link"`
	ItemTitle       string     `json:"item_title"`
	Updated         string     `json:"updated"`
	UpdatedParsed   *time.Time `json:"updated_parsed,omitempty"`
	Published       string     `json:"published,omitempty"`
	PublishedParsed *time.Time `json:"published_parsed,omitempty"`
	Author          string     `json:"author,omitempty"`
	Authors         []string   `json:"authors,omitempty"`
	GUID            string     `json:"guid,omitempty"`
	Categories      []string   `json:"categories,omitempty"`
}

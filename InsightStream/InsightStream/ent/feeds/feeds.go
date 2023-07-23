// Code generated by ent, DO NOT EDIT.

package entfeeds

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the baseFeeds type in the database.
	Label = "baseFeeds"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSiteURL holds the string denoting the site_url field in the database.
	FieldSiteURL = "site_url"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFeedURL holds the string denoting the feed_url field in the database.
	FieldFeedURL = "feed_url"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldDtCreated holds the string denoting the dt_created field in the database.
	FieldDtCreated = "dt_created"
	// FieldDtUpdated holds the string denoting the dt_updated field in the database.
	FieldDtUpdated = "dt_updated"
	// FieldFavorites holds the string denoting the favorites field in the database.
	FieldFavorites = "favorites"
	// Table holds the table name of the baseFeeds in the database.
	Table = "baseFeeds"
)

// Columns holds all SQL columns for baseFeeds fields.
var Columns = []string{
	FieldID,
	FieldSiteURL,
	FieldTitle,
	FieldDescription,
	FieldFeedURL,
	FieldLanguage,
	FieldDtCreated,
	FieldDtUpdated,
	FieldFavorites,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSiteURL holds the default value on creation for the "site_url" field.
	DefaultSiteURL string
	// SiteURLValidator is a validator for the "site_url" field. It is called by the builders before save.
	SiteURLValidator func(string) error
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultFeedURL holds the default value on creation for the "feed_url" field.
	DefaultFeedURL string
	// FeedURLValidator is a validator for the "feed_url" field. It is called by the builders before save.
	FeedURLValidator func(string) error
	// DefaultLanguage holds the default value on creation for the "language" field.
	DefaultLanguage string
	// LanguageValidator is a validator for the "language" field. It is called by the builders before save.
	LanguageValidator func(string) error
	// DefaultDtCreated holds the default value on creation for the "dt_created" field.
	DefaultDtCreated func() time.Time
	// UpdateDefaultDtUpdated holds the default value on update for the "dt_updated" field.
	UpdateDefaultDtUpdated func() time.Time
	// DefaultFavorites holds the default value on creation for the "favorites" field.
	DefaultFavorites int64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

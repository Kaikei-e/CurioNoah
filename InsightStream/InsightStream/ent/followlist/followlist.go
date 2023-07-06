// Code generated by ent, DO NOT EDIT.

package followlist

import (
	"insightstream/models/feeds"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the followlist type in the database.
	Label = "follow_list"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldXMLVersion holds the string denoting the xml_version field in the database.
	FieldXMLVersion = "xml_version"
	// FieldRssVersion holds the string denoting the rss_version field in the database.
	FieldRssVersion = "rss_version"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldLinks holds the string denoting the links field in the database.
	FieldLinks = "links"
	// FieldItemDescription holds the string denoting the item_description field in the database.
	FieldItemDescription = "item_description"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldDtCreated holds the string denoting the dt_created field in the database.
	FieldDtCreated = "dt_created"
	// FieldDtUpdated holds the string denoting the dt_updated field in the database.
	FieldDtUpdated = "dt_updated"
	// FieldDtLastInserted holds the string denoting the dt_last_inserted field in the database.
	FieldDtLastInserted = "dt_last_inserted"
	// FieldFeedCategory holds the string denoting the feed_category field in the database.
	FieldFeedCategory = "feed_category"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// FieldIsFavorite holds the string denoting the is_favorite field in the database.
	FieldIsFavorite = "is_favorite"
	// FieldIsRead holds the string denoting the is_read field in the database.
	FieldIsRead = "is_read"
	// FieldIsUpdated holds the string denoting the is_updated field in the database.
	FieldIsUpdated = "is_updated"
	// Table holds the table name of the followlist in the database.
	Table = "follow_lists"
)

// Columns holds all SQL columns for followlist fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldXMLVersion,
	FieldRssVersion,
	FieldURL,
	FieldTitle,
	FieldDescription,
	FieldLink,
	FieldLinks,
	FieldItemDescription,
	FieldLanguage,
	FieldDtCreated,
	FieldDtUpdated,
	FieldDtLastInserted,
	FieldFeedCategory,
	FieldIsActive,
	FieldIsFavorite,
	FieldIsRead,
	FieldIsUpdated,
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
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
	// DefaultXMLVersion holds the default value on creation for the "xml_version" field.
	DefaultXMLVersion int8
	// DefaultRssVersion holds the default value on creation for the "rss_version" field.
	DefaultRssVersion int8
	// DefaultURL holds the default value on creation for the "url" field.
	DefaultURL string
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultTitle holds the default value on creation for the "title" field.
	DefaultTitle string
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultLink holds the default value on creation for the "link" field.
	DefaultLink string
	// DefaultLinks holds the default value on creation for the "links" field.
	DefaultLinks feeds.FeedLink
	// DefaultLanguage holds the default value on creation for the "language" field.
	DefaultLanguage string
	// LanguageValidator is a validator for the "language" field. It is called by the builders before save.
	LanguageValidator func(string) error
	// DefaultDtCreated holds the default value on creation for the "dt_created" field.
	DefaultDtCreated time.Time
	// DefaultDtUpdated holds the default value on creation for the "dt_updated" field.
	DefaultDtUpdated time.Time
	// DefaultDtLastInserted holds the default value on creation for the "dt_last_inserted" field.
	DefaultDtLastInserted time.Time
	// DefaultFeedCategory holds the default value on creation for the "feed_category" field.
	DefaultFeedCategory int
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultIsFavorite holds the default value on creation for the "is_favorite" field.
	DefaultIsFavorite bool
	// DefaultIsRead holds the default value on creation for the "is_read" field.
	DefaultIsRead bool
	// DefaultIsUpdated holds the default value on creation for the "is_updated" field.
	DefaultIsUpdated bool
)

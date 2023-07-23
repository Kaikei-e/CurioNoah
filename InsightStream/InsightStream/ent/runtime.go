// Code generated by ent, DO NOT EDIT.

package ent

import (
	"insightstream/domain/baseFeeds"
	"insightstream/ent/cooccurrencenetworkpool"
	"insightstream/ent/feedaudittrailaction"
	"insightstream/ent/feedaudittraillog"
	entfeeds "insightstream/ent/feeds"
	"insightstream/ent/followlist"
	"insightstream/ent/schema"
	"insightstream/ent/users"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cooccurrencenetworkpoolFields := schema.CooccurrenceNetworkPool{}.Fields()
	_ = cooccurrencenetworkpoolFields
	// cooccurrencenetworkpoolDescSiteURL is the schema descriptor for site_url field.
	cooccurrencenetworkpoolDescSiteURL := cooccurrencenetworkpoolFields[1].Descriptor()
	// cooccurrencenetworkpool.SiteURLValidator is a validator for the "site_url" field. It is called by the builders before save.
	cooccurrencenetworkpool.SiteURLValidator = cooccurrencenetworkpoolDescSiteURL.Validators[0].(func(string) error)
	// cooccurrencenetworkpoolDescTitles is the schema descriptor for titles field.
	cooccurrencenetworkpoolDescTitles := cooccurrencenetworkpoolFields[2].Descriptor()
	// cooccurrencenetworkpool.DefaultTitles holds the default value on creation for the titles field.
	cooccurrencenetworkpool.DefaultTitles = cooccurrencenetworkpoolDescTitles.Default.([]string)
	// cooccurrencenetworkpoolDescDescriptions is the schema descriptor for descriptions field.
	cooccurrencenetworkpoolDescDescriptions := cooccurrencenetworkpoolFields[3].Descriptor()
	// cooccurrencenetworkpool.DefaultDescriptions holds the default value on creation for the descriptions field.
	cooccurrencenetworkpool.DefaultDescriptions = cooccurrencenetworkpoolDescDescriptions.Default.([]string)
	// cooccurrencenetworkpoolDescID is the schema descriptor for id field.
	cooccurrencenetworkpoolDescID := cooccurrencenetworkpoolFields[0].Descriptor()
	// cooccurrencenetworkpool.DefaultID holds the default value on creation for the id field.
	cooccurrencenetworkpool.DefaultID = cooccurrencenetworkpoolDescID.Default.(func() uuid.UUID)
	feedaudittrailactionFields := schema.FeedAuditTrailAction{}.Fields()
	_ = feedaudittrailactionFields
	// feedaudittrailactionDescAction is the schema descriptor for action field.
	feedaudittrailactionDescAction := feedaudittrailactionFields[1].Descriptor()
	// feedaudittrailaction.ActionValidator is a validator for the "action" field. It is called by the builders before save.
	feedaudittrailaction.ActionValidator = feedaudittrailactionDescAction.Validators[0].(func(string) error)
	feedaudittraillogFields := schema.FeedAuditTrailLog{}.Fields()
	_ = feedaudittraillogFields
	// feedaudittraillogDescUpdatedAt is the schema descriptor for updated_at field.
	feedaudittraillogDescUpdatedAt := feedaudittraillogFields[1].Descriptor()
	// feedaudittraillog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feedaudittraillog.DefaultUpdatedAt = feedaudittraillogDescUpdatedAt.Default.(func() time.Time)
	// feedaudittraillog.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feedaudittraillog.UpdateDefaultUpdatedAt = feedaudittraillogDescUpdatedAt.UpdateDefault.(func() time.Time)
	entfeedsFields := schema.Feeds{}.Fields()
	_ = entfeedsFields
	// entfeedsDescSiteURL is the schema descriptor for site_url field.
	entfeedsDescSiteURL := entfeedsFields[1].Descriptor()
	// entfeeds.DefaultSiteURL holds the default value on creation for the site_url field.
	entfeeds.DefaultSiteURL = entfeedsDescSiteURL.Default.(string)
	// entfeeds.SiteURLValidator is a validator for the "site_url" field. It is called by the builders before save.
	entfeeds.SiteURLValidator = entfeedsDescSiteURL.Validators[0].(func(string) error)
	// entfeedsDescTitle is the schema descriptor for title field.
	entfeedsDescTitle := entfeedsFields[2].Descriptor()
	// entfeeds.DefaultTitle holds the default value on creation for the title field.
	entfeeds.DefaultTitle = entfeedsDescTitle.Default.(string)
	// entfeeds.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	entfeeds.TitleValidator = entfeedsDescTitle.Validators[0].(func(string) error)
	// entfeedsDescDescription is the schema descriptor for description field.
	entfeedsDescDescription := entfeedsFields[3].Descriptor()
	// entfeeds.DefaultDescription holds the default value on creation for the description field.
	entfeeds.DefaultDescription = entfeedsDescDescription.Default.(string)
	// entfeeds.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	entfeeds.DescriptionValidator = entfeedsDescDescription.Validators[0].(func(string) error)
	// entfeedsDescFeedURL is the schema descriptor for feed_url field.
	entfeedsDescFeedURL := entfeedsFields[4].Descriptor()
	// entfeeds.DefaultFeedURL holds the default value on creation for the feed_url field.
	entfeeds.DefaultFeedURL = entfeedsDescFeedURL.Default.(string)
	// entfeeds.FeedURLValidator is a validator for the "feed_url" field. It is called by the builders before save.
	entfeeds.FeedURLValidator = entfeedsDescFeedURL.Validators[0].(func(string) error)
	// entfeedsDescLanguage is the schema descriptor for language field.
	entfeedsDescLanguage := entfeedsFields[5].Descriptor()
	// entfeeds.DefaultLanguage holds the default value on creation for the language field.
	entfeeds.DefaultLanguage = entfeedsDescLanguage.Default.(string)
	// entfeeds.LanguageValidator is a validator for the "language" field. It is called by the builders before save.
	entfeeds.LanguageValidator = entfeedsDescLanguage.Validators[0].(func(string) error)
	// entfeedsDescDtCreated is the schema descriptor for dt_created field.
	entfeedsDescDtCreated := entfeedsFields[6].Descriptor()
	// entfeeds.DefaultDtCreated holds the default value on creation for the dt_created field.
	entfeeds.DefaultDtCreated = entfeedsDescDtCreated.Default.(func() time.Time)
	// entfeedsDescDtUpdated is the schema descriptor for dt_updated field.
	entfeedsDescDtUpdated := entfeedsFields[7].Descriptor()
	// entfeeds.UpdateDefaultDtUpdated holds the default value on update for the dt_updated field.
	entfeeds.UpdateDefaultDtUpdated = entfeedsDescDtUpdated.UpdateDefault.(func() time.Time)
	// entfeedsDescFavorites is the schema descriptor for favorites field.
	entfeedsDescFavorites := entfeedsFields[8].Descriptor()
	// entfeeds.DefaultFavorites holds the default value on creation for the favorites field.
	entfeeds.DefaultFavorites = entfeedsDescFavorites.Default.(int64)
	// entfeedsDescID is the schema descriptor for id field.
	entfeedsDescID := entfeedsFields[0].Descriptor()
	// entfeeds.DefaultID holds the default value on creation for the id field.
	entfeeds.DefaultID = entfeedsDescID.Default.(func() uuid.UUID)
	followlistFields := schema.FollowList{}.Fields()
	_ = followlistFields
	// followlistDescUUID is the schema descriptor for uuid field.
	followlistDescUUID := followlistFields[0].Descriptor()
	// followlist.DefaultUUID holds the default value on creation for the uuid field.
	followlist.DefaultUUID = followlistDescUUID.Default.(func() uuid.UUID)
	// followlistDescXMLVersion is the schema descriptor for xml_version field.
	followlistDescXMLVersion := followlistFields[1].Descriptor()
	// followlist.DefaultXMLVersion holds the default value on creation for the xml_version field.
	followlist.DefaultXMLVersion = followlistDescXMLVersion.Default.(int8)
	// followlistDescRssVersion is the schema descriptor for rss_version field.
	followlistDescRssVersion := followlistFields[2].Descriptor()
	// followlist.DefaultRssVersion holds the default value on creation for the rss_version field.
	followlist.DefaultRssVersion = followlistDescRssVersion.Default.(int8)
	// followlistDescURL is the schema descriptor for url field.
	followlistDescURL := followlistFields[3].Descriptor()
	// followlist.DefaultURL holds the default value on creation for the url field.
	followlist.DefaultURL = followlistDescURL.Default.(string)
	// followlist.URLValidator is a validator for the "url" field. It is called by the builders before save.
	followlist.URLValidator = followlistDescURL.Validators[0].(func(string) error)
	// followlistDescTitle is the schema descriptor for title field.
	followlistDescTitle := followlistFields[4].Descriptor()
	// followlist.DefaultTitle holds the default value on creation for the title field.
	followlist.DefaultTitle = followlistDescTitle.Default.(string)
	// followlistDescDescription is the schema descriptor for description field.
	followlistDescDescription := followlistFields[5].Descriptor()
	// followlist.DefaultDescription holds the default value on creation for the description field.
	followlist.DefaultDescription = followlistDescDescription.Default.(string)
	// followlistDescLink is the schema descriptor for link field.
	followlistDescLink := followlistFields[6].Descriptor()
	// followlist.DefaultLink holds the default value on creation for the link field.
	followlist.DefaultLink = followlistDescLink.Default.(string)
	// followlistDescLinks is the schema descriptor for links field.
	followlistDescLinks := followlistFields[7].Descriptor()
	// followlist.DefaultLinks holds the default value on creation for the links field.
	followlist.DefaultLinks = followlistDescLinks.Default.(baseFeeds.FeedLink)
	// followlistDescLanguage is the schema descriptor for language field.
	followlistDescLanguage := followlistFields[9].Descriptor()
	// followlist.DefaultLanguage holds the default value on creation for the language field.
	followlist.DefaultLanguage = followlistDescLanguage.Default.(string)
	// followlist.LanguageValidator is a validator for the "language" field. It is called by the builders before save.
	followlist.LanguageValidator = followlistDescLanguage.Validators[0].(func(string) error)
	// followlistDescDtCreated is the schema descriptor for dt_created field.
	followlistDescDtCreated := followlistFields[10].Descriptor()
	// followlist.DefaultDtCreated holds the default value on creation for the dt_created field.
	followlist.DefaultDtCreated = followlistDescDtCreated.Default.(time.Time)
	// followlistDescDtUpdated is the schema descriptor for dt_updated field.
	followlistDescDtUpdated := followlistFields[11].Descriptor()
	// followlist.DefaultDtUpdated holds the default value on creation for the dt_updated field.
	followlist.DefaultDtUpdated = followlistDescDtUpdated.Default.(time.Time)
	// followlistDescDtLastInserted is the schema descriptor for dt_last_inserted field.
	followlistDescDtLastInserted := followlistFields[12].Descriptor()
	// followlist.DefaultDtLastInserted holds the default value on creation for the dt_last_inserted field.
	followlist.DefaultDtLastInserted = followlistDescDtLastInserted.Default.(time.Time)
	// followlistDescFeedCategory is the schema descriptor for feed_category field.
	followlistDescFeedCategory := followlistFields[13].Descriptor()
	// followlist.DefaultFeedCategory holds the default value on creation for the feed_category field.
	followlist.DefaultFeedCategory = followlistDescFeedCategory.Default.(int)
	// followlistDescIsActive is the schema descriptor for is_active field.
	followlistDescIsActive := followlistFields[14].Descriptor()
	// followlist.DefaultIsActive holds the default value on creation for the is_active field.
	followlist.DefaultIsActive = followlistDescIsActive.Default.(bool)
	// followlistDescIsFavorite is the schema descriptor for is_favorite field.
	followlistDescIsFavorite := followlistFields[15].Descriptor()
	// followlist.DefaultIsFavorite holds the default value on creation for the is_favorite field.
	followlist.DefaultIsFavorite = followlistDescIsFavorite.Default.(bool)
	// followlistDescIsRead is the schema descriptor for is_read field.
	followlistDescIsRead := followlistFields[16].Descriptor()
	// followlist.DefaultIsRead holds the default value on creation for the is_read field.
	followlist.DefaultIsRead = followlistDescIsRead.Default.(bool)
	// followlistDescIsUpdated is the schema descriptor for is_updated field.
	followlistDescIsUpdated := followlistFields[17].Descriptor()
	// followlist.DefaultIsUpdated holds the default value on creation for the is_updated field.
	followlist.DefaultIsUpdated = followlistDescIsUpdated.Default.(bool)
	usersFields := schema.Users{}.Fields()
	_ = usersFields
	// usersDescUsername is the schema descriptor for username field.
	usersDescUsername := usersFields[1].Descriptor()
	// users.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	users.UsernameValidator = usersDescUsername.Validators[0].(func(string) error)
	// usersDescHashedPassword is the schema descriptor for hashed_password field.
	usersDescHashedPassword := usersFields[2].Descriptor()
	// users.HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	users.HashedPasswordValidator = usersDescHashedPassword.Validators[0].(func([]byte) error)
	// usersDescCreatedAt is the schema descriptor for created_at field.
	usersDescCreatedAt := usersFields[3].Descriptor()
	// users.DefaultCreatedAt holds the default value on creation for the created_at field.
	users.DefaultCreatedAt = usersDescCreatedAt.Default.(time.Time)
	// usersDescUpdatedAt is the schema descriptor for updated_at field.
	usersDescUpdatedAt := usersFields[4].Descriptor()
	// users.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	users.DefaultUpdatedAt = usersDescUpdatedAt.Default.(time.Time)
}

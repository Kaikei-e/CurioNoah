// Code generated by ent, DO NOT EDIT.

package followlists

import (
	"insightstream/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldUUID, v))
}

// XMLVersion applies equality check predicate on the "xml_version" field. It's identical to XMLVersionEQ.
func XMLVersion(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldXMLVersion, v))
}

// RssVersion applies equality check predicate on the "rss_version" field. It's identical to RssVersionEQ.
func RssVersion(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldRssVersion, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldURL, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDescription, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldLink, v))
}

// Language applies equality check predicate on the "language" field. It's identical to LanguageEQ.
func Language(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldLanguage, v))
}

// DtCreated applies equality check predicate on the "dt_created" field. It's identical to DtCreatedEQ.
func DtCreated(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDtCreated, v))
}

// DtUpdated applies equality check predicate on the "dt_updated" field. It's identical to DtUpdatedEQ.
func DtUpdated(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDtUpdated, v))
}

// DtLastInserted applies equality check predicate on the "dt_last_inserted" field. It's identical to DtLastInsertedEQ.
func DtLastInserted(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDtLastInserted, v))
}

// FeedCategory applies equality check predicate on the "feed_category" field. It's identical to FeedCategoryEQ.
func FeedCategory(v int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldFeedCategory, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsActive, v))
}

// IsFavorite applies equality check predicate on the "is_favorite" field. It's identical to IsFavoriteEQ.
func IsFavorite(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsFavorite, v))
}

// IsRead applies equality check predicate on the "is_read" field. It's identical to IsReadEQ.
func IsRead(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsRead, v))
}

// IsUpdated applies equality check predicate on the "is_updated" field. It's identical to IsUpdatedEQ.
func IsUpdated(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsUpdated, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldUUID, v))
}

// XMLVersionEQ applies the EQ predicate on the "xml_version" field.
func XMLVersionEQ(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldXMLVersion, v))
}

// XMLVersionNEQ applies the NEQ predicate on the "xml_version" field.
func XMLVersionNEQ(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldXMLVersion, v))
}

// XMLVersionIn applies the In predicate on the "xml_version" field.
func XMLVersionIn(vs ...int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldXMLVersion, vs...))
}

// XMLVersionNotIn applies the NotIn predicate on the "xml_version" field.
func XMLVersionNotIn(vs ...int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldXMLVersion, vs...))
}

// XMLVersionGT applies the GT predicate on the "xml_version" field.
func XMLVersionGT(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldXMLVersion, v))
}

// XMLVersionGTE applies the GTE predicate on the "xml_version" field.
func XMLVersionGTE(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldXMLVersion, v))
}

// XMLVersionLT applies the LT predicate on the "xml_version" field.
func XMLVersionLT(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldXMLVersion, v))
}

// XMLVersionLTE applies the LTE predicate on the "xml_version" field.
func XMLVersionLTE(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldXMLVersion, v))
}

// RssVersionEQ applies the EQ predicate on the "rss_version" field.
func RssVersionEQ(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldRssVersion, v))
}

// RssVersionNEQ applies the NEQ predicate on the "rss_version" field.
func RssVersionNEQ(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldRssVersion, v))
}

// RssVersionIn applies the In predicate on the "rss_version" field.
func RssVersionIn(vs ...int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldRssVersion, vs...))
}

// RssVersionNotIn applies the NotIn predicate on the "rss_version" field.
func RssVersionNotIn(vs ...int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldRssVersion, vs...))
}

// RssVersionGT applies the GT predicate on the "rss_version" field.
func RssVersionGT(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldRssVersion, v))
}

// RssVersionGTE applies the GTE predicate on the "rss_version" field.
func RssVersionGTE(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldRssVersion, v))
}

// RssVersionLT applies the LT predicate on the "rss_version" field.
func RssVersionLT(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldRssVersion, v))
}

// RssVersionLTE applies the LTE predicate on the "rss_version" field.
func RssVersionLTE(v int8) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldRssVersion, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContainsFold(FieldURL, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContainsFold(FieldDescription, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContainsFold(FieldLink, v))
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldLanguage, v))
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldLanguage, v))
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldLanguage, vs...))
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldLanguage, vs...))
}

// LanguageGT applies the GT predicate on the "language" field.
func LanguageGT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldLanguage, v))
}

// LanguageGTE applies the GTE predicate on the "language" field.
func LanguageGTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldLanguage, v))
}

// LanguageLT applies the LT predicate on the "language" field.
func LanguageLT(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldLanguage, v))
}

// LanguageLTE applies the LTE predicate on the "language" field.
func LanguageLTE(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldLanguage, v))
}

// LanguageContains applies the Contains predicate on the "language" field.
func LanguageContains(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContains(FieldLanguage, v))
}

// LanguageHasPrefix applies the HasPrefix predicate on the "language" field.
func LanguageHasPrefix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasPrefix(FieldLanguage, v))
}

// LanguageHasSuffix applies the HasSuffix predicate on the "language" field.
func LanguageHasSuffix(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldHasSuffix(FieldLanguage, v))
}

// LanguageEqualFold applies the EqualFold predicate on the "language" field.
func LanguageEqualFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEqualFold(FieldLanguage, v))
}

// LanguageContainsFold applies the ContainsFold predicate on the "language" field.
func LanguageContainsFold(v string) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldContainsFold(FieldLanguage, v))
}

// DtCreatedEQ applies the EQ predicate on the "dt_created" field.
func DtCreatedEQ(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDtCreated, v))
}

// DtCreatedNEQ applies the NEQ predicate on the "dt_created" field.
func DtCreatedNEQ(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldDtCreated, v))
}

// DtCreatedIn applies the In predicate on the "dt_created" field.
func DtCreatedIn(vs ...time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldDtCreated, vs...))
}

// DtCreatedNotIn applies the NotIn predicate on the "dt_created" field.
func DtCreatedNotIn(vs ...time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldDtCreated, vs...))
}

// DtCreatedGT applies the GT predicate on the "dt_created" field.
func DtCreatedGT(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldDtCreated, v))
}

// DtCreatedGTE applies the GTE predicate on the "dt_created" field.
func DtCreatedGTE(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldDtCreated, v))
}

// DtCreatedLT applies the LT predicate on the "dt_created" field.
func DtCreatedLT(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldDtCreated, v))
}

// DtCreatedLTE applies the LTE predicate on the "dt_created" field.
func DtCreatedLTE(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldDtCreated, v))
}

// DtCreatedIsNil applies the IsNil predicate on the "dt_created" field.
func DtCreatedIsNil() predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIsNull(FieldDtCreated))
}

// DtCreatedNotNil applies the NotNil predicate on the "dt_created" field.
func DtCreatedNotNil() predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotNull(FieldDtCreated))
}

// DtUpdatedEQ applies the EQ predicate on the "dt_updated" field.
func DtUpdatedEQ(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDtUpdated, v))
}

// DtUpdatedNEQ applies the NEQ predicate on the "dt_updated" field.
func DtUpdatedNEQ(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldDtUpdated, v))
}

// DtUpdatedIn applies the In predicate on the "dt_updated" field.
func DtUpdatedIn(vs ...time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldDtUpdated, vs...))
}

// DtUpdatedNotIn applies the NotIn predicate on the "dt_updated" field.
func DtUpdatedNotIn(vs ...time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldDtUpdated, vs...))
}

// DtUpdatedGT applies the GT predicate on the "dt_updated" field.
func DtUpdatedGT(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldDtUpdated, v))
}

// DtUpdatedGTE applies the GTE predicate on the "dt_updated" field.
func DtUpdatedGTE(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldDtUpdated, v))
}

// DtUpdatedLT applies the LT predicate on the "dt_updated" field.
func DtUpdatedLT(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldDtUpdated, v))
}

// DtUpdatedLTE applies the LTE predicate on the "dt_updated" field.
func DtUpdatedLTE(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldDtUpdated, v))
}

// DtUpdatedIsNil applies the IsNil predicate on the "dt_updated" field.
func DtUpdatedIsNil() predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIsNull(FieldDtUpdated))
}

// DtUpdatedNotNil applies the NotNil predicate on the "dt_updated" field.
func DtUpdatedNotNil() predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotNull(FieldDtUpdated))
}

// DtLastInsertedEQ applies the EQ predicate on the "dt_last_inserted" field.
func DtLastInsertedEQ(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldDtLastInserted, v))
}

// DtLastInsertedNEQ applies the NEQ predicate on the "dt_last_inserted" field.
func DtLastInsertedNEQ(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldDtLastInserted, v))
}

// DtLastInsertedIn applies the In predicate on the "dt_last_inserted" field.
func DtLastInsertedIn(vs ...time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldDtLastInserted, vs...))
}

// DtLastInsertedNotIn applies the NotIn predicate on the "dt_last_inserted" field.
func DtLastInsertedNotIn(vs ...time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldDtLastInserted, vs...))
}

// DtLastInsertedGT applies the GT predicate on the "dt_last_inserted" field.
func DtLastInsertedGT(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldDtLastInserted, v))
}

// DtLastInsertedGTE applies the GTE predicate on the "dt_last_inserted" field.
func DtLastInsertedGTE(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldDtLastInserted, v))
}

// DtLastInsertedLT applies the LT predicate on the "dt_last_inserted" field.
func DtLastInsertedLT(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldDtLastInserted, v))
}

// DtLastInsertedLTE applies the LTE predicate on the "dt_last_inserted" field.
func DtLastInsertedLTE(v time.Time) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldDtLastInserted, v))
}

// DtLastInsertedIsNil applies the IsNil predicate on the "dt_last_inserted" field.
func DtLastInsertedIsNil() predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIsNull(FieldDtLastInserted))
}

// DtLastInsertedNotNil applies the NotNil predicate on the "dt_last_inserted" field.
func DtLastInsertedNotNil() predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotNull(FieldDtLastInserted))
}

// FeedCategoryEQ applies the EQ predicate on the "feed_category" field.
func FeedCategoryEQ(v int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldFeedCategory, v))
}

// FeedCategoryNEQ applies the NEQ predicate on the "feed_category" field.
func FeedCategoryNEQ(v int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldFeedCategory, v))
}

// FeedCategoryIn applies the In predicate on the "feed_category" field.
func FeedCategoryIn(vs ...int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldIn(FieldFeedCategory, vs...))
}

// FeedCategoryNotIn applies the NotIn predicate on the "feed_category" field.
func FeedCategoryNotIn(vs ...int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNotIn(FieldFeedCategory, vs...))
}

// FeedCategoryGT applies the GT predicate on the "feed_category" field.
func FeedCategoryGT(v int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGT(FieldFeedCategory, v))
}

// FeedCategoryGTE applies the GTE predicate on the "feed_category" field.
func FeedCategoryGTE(v int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldGTE(FieldFeedCategory, v))
}

// FeedCategoryLT applies the LT predicate on the "feed_category" field.
func FeedCategoryLT(v int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLT(FieldFeedCategory, v))
}

// FeedCategoryLTE applies the LTE predicate on the "feed_category" field.
func FeedCategoryLTE(v int) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldLTE(FieldFeedCategory, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldIsActive, v))
}

// IsFavoriteEQ applies the EQ predicate on the "is_favorite" field.
func IsFavoriteEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsFavorite, v))
}

// IsFavoriteNEQ applies the NEQ predicate on the "is_favorite" field.
func IsFavoriteNEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldIsFavorite, v))
}

// IsReadEQ applies the EQ predicate on the "is_read" field.
func IsReadEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsRead, v))
}

// IsReadNEQ applies the NEQ predicate on the "is_read" field.
func IsReadNEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldIsRead, v))
}

// IsUpdatedEQ applies the EQ predicate on the "is_updated" field.
func IsUpdatedEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldEQ(FieldIsUpdated, v))
}

// IsUpdatedNEQ applies the NEQ predicate on the "is_updated" field.
func IsUpdatedNEQ(v bool) predicate.FollowLists {
	return predicate.FollowLists(sql.FieldNEQ(FieldIsUpdated, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FollowLists) predicate.FollowLists {
	return predicate.FollowLists(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FollowLists) predicate.FollowLists {
	return predicate.FollowLists(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FollowLists) predicate.FollowLists {
	return predicate.FollowLists(sql.NotPredicates(p))
}

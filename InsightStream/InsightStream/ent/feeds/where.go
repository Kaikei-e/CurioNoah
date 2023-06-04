// Code generated by ent, DO NOT EDIT.

package entfeeds

import (
	"insightstream/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldID, id))
}

// SiteURL applies equality check predicate on the "site_url" field. It's identical to SiteURLEQ.
func SiteURL(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldSiteURL, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldDescription, v))
}

// FeedURL applies equality check predicate on the "feed_url" field. It's identical to FeedURLEQ.
func FeedURL(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldFeedURL, v))
}

// Language applies equality check predicate on the "language" field. It's identical to LanguageEQ.
func Language(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldLanguage, v))
}

// DtCreated applies equality check predicate on the "dt_created" field. It's identical to DtCreatedEQ.
func DtCreated(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldDtCreated, v))
}

// DtUpdated applies equality check predicate on the "dt_updated" field. It's identical to DtUpdatedEQ.
func DtUpdated(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldDtUpdated, v))
}

// Favorites applies equality check predicate on the "favorites" field. It's identical to FavoritesEQ.
func Favorites(v int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldFavorites, v))
}

// SiteURLEQ applies the EQ predicate on the "site_url" field.
func SiteURLEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldSiteURL, v))
}

// SiteURLNEQ applies the NEQ predicate on the "site_url" field.
func SiteURLNEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldSiteURL, v))
}

// SiteURLIn applies the In predicate on the "site_url" field.
func SiteURLIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldSiteURL, vs...))
}

// SiteURLNotIn applies the NotIn predicate on the "site_url" field.
func SiteURLNotIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldSiteURL, vs...))
}

// SiteURLGT applies the GT predicate on the "site_url" field.
func SiteURLGT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldSiteURL, v))
}

// SiteURLGTE applies the GTE predicate on the "site_url" field.
func SiteURLGTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldSiteURL, v))
}

// SiteURLLT applies the LT predicate on the "site_url" field.
func SiteURLLT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldSiteURL, v))
}

// SiteURLLTE applies the LTE predicate on the "site_url" field.
func SiteURLLTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldSiteURL, v))
}

// SiteURLContains applies the Contains predicate on the "site_url" field.
func SiteURLContains(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContains(FieldSiteURL, v))
}

// SiteURLHasPrefix applies the HasPrefix predicate on the "site_url" field.
func SiteURLHasPrefix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasPrefix(FieldSiteURL, v))
}

// SiteURLHasSuffix applies the HasSuffix predicate on the "site_url" field.
func SiteURLHasSuffix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasSuffix(FieldSiteURL, v))
}

// SiteURLEqualFold applies the EqualFold predicate on the "site_url" field.
func SiteURLEqualFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEqualFold(FieldSiteURL, v))
}

// SiteURLContainsFold applies the ContainsFold predicate on the "site_url" field.
func SiteURLContainsFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContainsFold(FieldSiteURL, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContainsFold(FieldDescription, v))
}

// FeedURLEQ applies the EQ predicate on the "feed_url" field.
func FeedURLEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldFeedURL, v))
}

// FeedURLNEQ applies the NEQ predicate on the "feed_url" field.
func FeedURLNEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldFeedURL, v))
}

// FeedURLIn applies the In predicate on the "feed_url" field.
func FeedURLIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldFeedURL, vs...))
}

// FeedURLNotIn applies the NotIn predicate on the "feed_url" field.
func FeedURLNotIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldFeedURL, vs...))
}

// FeedURLGT applies the GT predicate on the "feed_url" field.
func FeedURLGT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldFeedURL, v))
}

// FeedURLGTE applies the GTE predicate on the "feed_url" field.
func FeedURLGTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldFeedURL, v))
}

// FeedURLLT applies the LT predicate on the "feed_url" field.
func FeedURLLT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldFeedURL, v))
}

// FeedURLLTE applies the LTE predicate on the "feed_url" field.
func FeedURLLTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldFeedURL, v))
}

// FeedURLContains applies the Contains predicate on the "feed_url" field.
func FeedURLContains(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContains(FieldFeedURL, v))
}

// FeedURLHasPrefix applies the HasPrefix predicate on the "feed_url" field.
func FeedURLHasPrefix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasPrefix(FieldFeedURL, v))
}

// FeedURLHasSuffix applies the HasSuffix predicate on the "feed_url" field.
func FeedURLHasSuffix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasSuffix(FieldFeedURL, v))
}

// FeedURLEqualFold applies the EqualFold predicate on the "feed_url" field.
func FeedURLEqualFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEqualFold(FieldFeedURL, v))
}

// FeedURLContainsFold applies the ContainsFold predicate on the "feed_url" field.
func FeedURLContainsFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContainsFold(FieldFeedURL, v))
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldLanguage, v))
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldLanguage, v))
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldLanguage, vs...))
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...string) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldLanguage, vs...))
}

// LanguageGT applies the GT predicate on the "language" field.
func LanguageGT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldLanguage, v))
}

// LanguageGTE applies the GTE predicate on the "language" field.
func LanguageGTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldLanguage, v))
}

// LanguageLT applies the LT predicate on the "language" field.
func LanguageLT(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldLanguage, v))
}

// LanguageLTE applies the LTE predicate on the "language" field.
func LanguageLTE(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldLanguage, v))
}

// LanguageContains applies the Contains predicate on the "language" field.
func LanguageContains(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContains(FieldLanguage, v))
}

// LanguageHasPrefix applies the HasPrefix predicate on the "language" field.
func LanguageHasPrefix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasPrefix(FieldLanguage, v))
}

// LanguageHasSuffix applies the HasSuffix predicate on the "language" field.
func LanguageHasSuffix(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldHasSuffix(FieldLanguage, v))
}

// LanguageEqualFold applies the EqualFold predicate on the "language" field.
func LanguageEqualFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldEqualFold(FieldLanguage, v))
}

// LanguageContainsFold applies the ContainsFold predicate on the "language" field.
func LanguageContainsFold(v string) predicate.Feeds {
	return predicate.Feeds(sql.FieldContainsFold(FieldLanguage, v))
}

// DtCreatedEQ applies the EQ predicate on the "dt_created" field.
func DtCreatedEQ(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldDtCreated, v))
}

// DtCreatedNEQ applies the NEQ predicate on the "dt_created" field.
func DtCreatedNEQ(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldDtCreated, v))
}

// DtCreatedIn applies the In predicate on the "dt_created" field.
func DtCreatedIn(vs ...time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldDtCreated, vs...))
}

// DtCreatedNotIn applies the NotIn predicate on the "dt_created" field.
func DtCreatedNotIn(vs ...time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldDtCreated, vs...))
}

// DtCreatedGT applies the GT predicate on the "dt_created" field.
func DtCreatedGT(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldDtCreated, v))
}

// DtCreatedGTE applies the GTE predicate on the "dt_created" field.
func DtCreatedGTE(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldDtCreated, v))
}

// DtCreatedLT applies the LT predicate on the "dt_created" field.
func DtCreatedLT(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldDtCreated, v))
}

// DtCreatedLTE applies the LTE predicate on the "dt_created" field.
func DtCreatedLTE(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldDtCreated, v))
}

// DtUpdatedEQ applies the EQ predicate on the "dt_updated" field.
func DtUpdatedEQ(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldDtUpdated, v))
}

// DtUpdatedNEQ applies the NEQ predicate on the "dt_updated" field.
func DtUpdatedNEQ(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldDtUpdated, v))
}

// DtUpdatedIn applies the In predicate on the "dt_updated" field.
func DtUpdatedIn(vs ...time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldDtUpdated, vs...))
}

// DtUpdatedNotIn applies the NotIn predicate on the "dt_updated" field.
func DtUpdatedNotIn(vs ...time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldDtUpdated, vs...))
}

// DtUpdatedGT applies the GT predicate on the "dt_updated" field.
func DtUpdatedGT(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldDtUpdated, v))
}

// DtUpdatedGTE applies the GTE predicate on the "dt_updated" field.
func DtUpdatedGTE(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldDtUpdated, v))
}

// DtUpdatedLT applies the LT predicate on the "dt_updated" field.
func DtUpdatedLT(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldDtUpdated, v))
}

// DtUpdatedLTE applies the LTE predicate on the "dt_updated" field.
func DtUpdatedLTE(v time.Time) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldDtUpdated, v))
}

// FavoritesEQ applies the EQ predicate on the "favorites" field.
func FavoritesEQ(v int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldEQ(FieldFavorites, v))
}

// FavoritesNEQ applies the NEQ predicate on the "favorites" field.
func FavoritesNEQ(v int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldNEQ(FieldFavorites, v))
}

// FavoritesIn applies the In predicate on the "favorites" field.
func FavoritesIn(vs ...int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldIn(FieldFavorites, vs...))
}

// FavoritesNotIn applies the NotIn predicate on the "favorites" field.
func FavoritesNotIn(vs ...int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldNotIn(FieldFavorites, vs...))
}

// FavoritesGT applies the GT predicate on the "favorites" field.
func FavoritesGT(v int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldGT(FieldFavorites, v))
}

// FavoritesGTE applies the GTE predicate on the "favorites" field.
func FavoritesGTE(v int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldGTE(FieldFavorites, v))
}

// FavoritesLT applies the LT predicate on the "favorites" field.
func FavoritesLT(v int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldLT(FieldFavorites, v))
}

// FavoritesLTE applies the LTE predicate on the "favorites" field.
func FavoritesLTE(v int64) predicate.Feeds {
	return predicate.Feeds(sql.FieldLTE(FieldFavorites, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Feeds) predicate.Feeds {
	return predicate.Feeds(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Feeds) predicate.Feeds {
	return predicate.Feeds(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Feeds) predicate.Feeds {
	return predicate.Feeds(func(s *sql.Selector) {
		p(s.Not())
	})
}
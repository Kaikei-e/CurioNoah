// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	entfeeds "insightstream/ent/feeds"
)

// Feeds is the model entity for the Feeds schema.
type Feeds struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SiteURL holds the value of the "site_url" field.
	SiteURL string `json:"site_url,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// FeedURL holds the value of the "feed_url" field.
	FeedURL string `json:"feed_url,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// DtCreated holds the value of the "dt_created" field.
	DtCreated time.Time `json:"dt_created,omitempty"`
	// DtUpdated holds the value of the "dt_updated" field.
	DtUpdated time.Time `json:"dt_updated,omitempty"`
	// Favorites holds the value of the "favorites" field.
	Favorites int64 `json:"favorites,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Feeds) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entfeeds.FieldFavorites:
			values[i] = new(sql.NullInt64)
		case entfeeds.FieldSiteURL, entfeeds.FieldTitle, entfeeds.FieldDescription, entfeeds.FieldFeedURL, entfeeds.FieldLanguage:
			values[i] = new(sql.NullString)
		case entfeeds.FieldDtCreated, entfeeds.FieldDtUpdated:
			values[i] = new(sql.NullTime)
		case entfeeds.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Feeds", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Feeds fields.
func (f *Feeds) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entfeeds.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case entfeeds.FieldSiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field site_url", values[i])
			} else if value.Valid {
				f.SiteURL = value.String
			}
		case entfeeds.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				f.Title = value.String
			}
		case entfeeds.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				f.Description = value.String
			}
		case entfeeds.FieldFeedURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_url", values[i])
			} else if value.Valid {
				f.FeedURL = value.String
			}
		case entfeeds.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				f.Language = value.String
			}
		case entfeeds.FieldDtCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dt_created", values[i])
			} else if value.Valid {
				f.DtCreated = value.Time
			}
		case entfeeds.FieldDtUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dt_updated", values[i])
			} else if value.Valid {
				f.DtUpdated = value.Time
			}
		case entfeeds.FieldFavorites:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field favorites", values[i])
			} else if value.Valid {
				f.Favorites = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Feeds.
// Note that you need to call Feeds.Unwrap() before calling this method if this Feeds
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Feeds) Update() *FeedsUpdateOne {
	return (&FeedsClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Feeds entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Feeds) Unwrap() *Feeds {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Feeds is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Feeds) String() string {
	var builder strings.Builder
	builder.WriteString("Feeds(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("site_url=")
	builder.WriteString(f.SiteURL)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(f.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(f.Description)
	builder.WriteString(", ")
	builder.WriteString("feed_url=")
	builder.WriteString(f.FeedURL)
	builder.WriteString(", ")
	builder.WriteString("language=")
	builder.WriteString(f.Language)
	builder.WriteString(", ")
	builder.WriteString("dt_created=")
	builder.WriteString(f.DtCreated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("dt_updated=")
	builder.WriteString(f.DtUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("favorites=")
	builder.WriteString(fmt.Sprintf("%v", f.Favorites))
	builder.WriteByte(')')
	return builder.String()
}

// FeedsSlice is a parsable slice of Feeds.
type FeedsSlice []*Feeds

func (f FeedsSlice) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
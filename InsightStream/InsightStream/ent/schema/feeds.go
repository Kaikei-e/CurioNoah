package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Feeds holds the schema definition for the Feeds entity.
type Feeds struct {
	ent.Schema
}

// Fields of the Feeds.
func (Feeds) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique(),
		field.String("site_url").Default("").NotEmpty(),
		field.String("title").Default("").NotEmpty(),
		field.Text("description").Default("").NotEmpty(),
		field.String("feed_url").Default("").NotEmpty(),
		field.String("language").Default("").NotEmpty(),
		field.Time("dt_created").Immutable().Default(time.Now).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("dt_updated").UpdateDefault(time.Now).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Int64("favorites").Default(0),
	}
}

// Indexes of the Feeds.
func (Feeds) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		// unique index
		index.Fields("site_url", "feed_url").Unique(),
		index.Fields("id", "feed_url").Unique(),
		index.Fields("id", "dt_updated").Unique(),
		index.Fields("dt_updated", "feed_url").
			Annotations(entsql.Desc()).
			StorageKey("idx_feeds_dt_updated_feed_url"),
		index.Fields("title").Annotations(
			entsql.IndexTypes(map[string]string{
				dialect.MySQL:    "FULLTEXT",
				dialect.Postgres: "GIN",
			}),
		),
		index.Fields("description").Annotations(
			entsql.IndexTypes(map[string]string{
				dialect.MySQL:    "FULLTEXT",
				dialect.Postgres: "GIN",
			}),
		),
	}
}

// Edges of the Feeds.
func (Feeds) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"insightstream/domain/baseFeeds"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// FollowLists holds the schema definition for the FollowLists entity.
type FollowLists struct {
	ent.Schema
}

// Fields of the FollowLists.
func (FollowLists) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Int8("xml_version").Default(1),
		field.Int8("rss_version").Default(2),
		field.String("url").Default("").NotEmpty(),
		field.String("title").Default(""),
		field.String("description").Default(""),
		field.String("link").Default(""),
		field.JSON("links", baseFeeds.FeedLink{}).Default(baseFeeds.FeedLink{Link: []string{"{}"}}),
		field.JSON("item_description", []baseFeeds.FeedItem{}),
		field.String("language").Default("").NotEmpty(),
		field.Time("dt_created").Default(time.Now()).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("dt_updated").Default(time.Now()).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("dt_last_inserted").Default(time.Now()).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Int("feed_category").Default(0),
		field.Bool("is_active").Default(true),
		field.Bool("is_favorite").Default(false),
		field.Bool("is_read").Default(false),
		field.Bool("is_updated").Default(false),
	}
}

// Indexes of the FollowLists.
func (FollowLists) Indexes() []ent.Index {
	return []ent.Index{
		// unique index
		index.Fields("uuid").Unique(),
	}
}

// Edges of the FollowLists.
func (FollowLists) Edges() []ent.Edge {
	return nil
}

func (FollowLists) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "follow_lists"},
	}
}

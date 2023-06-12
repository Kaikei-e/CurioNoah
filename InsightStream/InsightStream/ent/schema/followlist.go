package schema

import (
	"entgo.io/ent/schema/index"
	"insightstream/models/feeds"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FollowList holds the schema definition for the FollowList entity.
type FollowList struct {
	ent.Schema
}

// Fields of the FollowList.
func (FollowList) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
		field.Int8("xml_version").Default(1),
		field.Int8("rss_version").Default(2),
		field.String("url"),
		field.String("title"),
		field.String("description"),
		field.String("link"),
		field.JSON("links", feeds.FeedLink{}),
		field.JSON("item_description", []feeds.FeedItem{}),
		field.String("language"),
		field.Time("dt_created").Default(time.Now()),
		field.Time("dt_updated").Default(time.Now()),
		field.Time("dt_last_inserted").Default(time.Now()),
		field.Int("feed_category").Default(0),
		field.Bool("is_active").Default(true),
		field.Bool("is_favorite").Default(false),
		field.Bool("is_read").Default(false),
		field.Bool("is_updated").Default(false),
	}
}

// Indexes of the FollowList.
func (FollowList) Indexes() []ent.Index {
	return []ent.Index{
		// unique index
		index.Fields("uuid").Unique(),
	}
}

// Edges of the FollowList.
func (FollowList) Edges() []ent.Edge {
	return nil
}

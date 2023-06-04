package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		field.String("description").Default("").NotEmpty(),
		field.String("feed_url").Default("").NotEmpty(),
		field.String("language").Default("").NotEmpty(),
		field.Time("dt_created").Immutable().Default(time.Now),
		field.Time("dt_updated").UpdateDefault(time.Now),
		field.Int64("favorites").Default(0),
	}
}

// Edges of the Feeds.
func (Feeds) Edges() []ent.Edge {
	return nil
}
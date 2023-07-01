package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type FeedAuditTrail struct {
	ent.Schema
}

// Fields of the FeedAuditTrail.
func (FeedAuditTrail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("action").Default("").NotEmpty(),
	}
}

// Indexes of the FeedAuditTrail.
func (FeedAuditTrail) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

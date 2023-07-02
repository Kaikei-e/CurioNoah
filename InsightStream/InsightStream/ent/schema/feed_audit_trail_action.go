package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type FeedAuditTrailAction struct {
	ent.Schema
}

// Fields of the FeedAuditTrailLog.
func (FeedAuditTrailAction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("action").Unique().NotEmpty(),
	}
}

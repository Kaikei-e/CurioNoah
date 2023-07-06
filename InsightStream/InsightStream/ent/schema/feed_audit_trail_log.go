package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type FeedAuditTrailLog struct {
	ent.Schema
}

// Fields of the FeedAuditTrailLog.
func (FeedAuditTrailLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Indexes of the FeedAuditTrailLog.
func (FeedAuditTrailLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// Edges of the FeedAuditTrailLog.
func (FeedAuditTrailLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("action", FeedAuditTrailAction.Type).
			StorageKey(edge.Column("action_id")).
			Unique().
			Required(),
	}
}

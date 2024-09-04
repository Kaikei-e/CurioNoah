package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type CooccurrenceNetworkPool struct {
	ent.Schema
}

func (CooccurrenceNetworkPool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "cooccurrence_network_pool"},
	}
}

func (CooccurrenceNetworkPool) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Unique(),
		field.String("site_url").NotEmpty().Unique(),
		field.JSON("titles", []string{}).Default([]string{}),
		field.JSON("descriptions", []string{}).Default([]string{}),
	}
}

func (CooccurrenceNetworkPool) Indexes() []ent.Index {
	return []ent.Index{
		// primary index
		index.Fields("id").Unique(),
		// unique index
		index.Fields("id", "site_url").Unique(),
	}
}

func (CooccurrenceNetworkPool) Edges() []ent.Edge {
	return nil
}

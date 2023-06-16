package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("username").NotEmpty(),
		field.Bytes("hashed_password").NotEmpty(),
	}
}

// Indexes of the Users.
func (Users) Indexes() []ent.Index {
	return []ent.Index{
		// unique index
		index.Fields("id").Unique(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}

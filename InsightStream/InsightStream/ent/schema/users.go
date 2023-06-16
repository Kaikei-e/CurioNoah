package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.Bytes("password").Sensitive().NotEmpty(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}

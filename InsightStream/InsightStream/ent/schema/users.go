package schema

import (
	"time"

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
		field.Bytes("hashed_password").Sensitive().NotEmpty(),
		field.Time("created_at").Immutable().Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

// Indexes of the Users.
func (Users) Indexes() []ent.Index {
	return []ent.Index{
		// unique index
		index.Fields("id").Unique(),
		index.Fields("username").Unique(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}

package schema

import "entgo.io/ent"

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return nil
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}

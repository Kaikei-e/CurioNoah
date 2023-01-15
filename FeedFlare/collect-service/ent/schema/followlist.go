package schema

import "entgo.io/ent"

// FollowList holds the schema definition for the FollowList entity.
type FollowList struct {
	ent.Schema
}

// Fields of the FollowList.
func (FollowList) Fields() []ent.Field {
	return nil
}

// Edges of the FollowList.
func (FollowList) Edges() []ent.Edge {
	return nil
}

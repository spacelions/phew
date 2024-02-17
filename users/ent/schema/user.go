package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique(),
		field.Int("age").
			Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Indexes of the User
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username").
			Unique(),
	}
}

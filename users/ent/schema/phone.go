package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Phone holds the schema definition for the Phone entity.
type Phone struct {
	ent.Schema
}

// Fields of the Phone.
func (Phone) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			NotEmpty(),
		field.String("country_code").
			NotEmpty().
			Default("+86"),
	}
}

// Edges of the Phone.
func (Phone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique(),
	}
}

// Indexes of the Phone
func (Phone) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country_code", "number").
			Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
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
		field.String("country_code").
			NotEmpty(),
		field.String("number").
			NotEmpty(),
	}
}

// Edges of the Phone.
func (Phone) Edges() []ent.Edge {
	return nil
}

func (Phone) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("country_code", "number").
			Unique(),
	}
}
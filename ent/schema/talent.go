package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Talent holds the schema definition for the Talent entity.
type Talent struct {
	ent.Schema
}

// Fields of the Talent.
func (Talent) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Talent.
func (Talent) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

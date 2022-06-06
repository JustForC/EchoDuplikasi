package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Nillable(),
		field.String("period").Optional().Nillable(),
		field.String("position").Optional().Nillable(),
		field.String("detail").Optional().Nillable(),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

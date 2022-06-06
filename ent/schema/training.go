package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Training holds the schema definition for the Training entity.
type Training struct {
	ent.Schema
}

// Fields of the Training.
func (Training) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Nillable(),
		field.String("period").Optional().Nillable(),
		field.String("year").Optional().Nillable(),
		field.String("organizer").Optional().Nillable(),
		field.String("certificate").Optional().Nillable(),
	}
}

// Edges of the Training.
func (Training) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

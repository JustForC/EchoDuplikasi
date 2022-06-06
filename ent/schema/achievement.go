package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Achievement holds the schema definition for the Achievement entity.
type Achievement struct {
	ent.Schema
}

// Fields of the Achievement.
func (Achievement) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Nillable(),
		field.String("organizer").Optional().Nillable(),
		field.String("level").Optional().Nillable(),
	}
}

// Edges of the Achievement.
func (Achievement) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

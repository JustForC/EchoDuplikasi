package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Nillable(),
		field.String("talk").Optional().Nillable(),
		field.String("write").Optional().Nillable(),
		field.String("read").Optional().Nillable(),
		field.String("listen").Optional().Nillable(),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

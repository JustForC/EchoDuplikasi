package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Education holds the schema definition for the Education entity.
type Education struct {
	ent.Schema
}

// Fields of the Education.
func (Education) Fields() []ent.Field {
	return []ent.Field{
		field.String("grade"),
		field.String("name"),
		field.String("province"),
		field.String("city"),
		field.String("enter"),
		field.String("graduate"),
	}
}

// Edges of the Education.
func (Education) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

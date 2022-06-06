package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Family holds the schema definition for the Family entity.
type Family struct {
	ent.Schema
}

// Fields of the Family.
func (Family) Fields() []ent.Field {
	return []ent.Field{
		field.String("status"),
		field.String("name"),
		field.String("gender"),
		field.String("birthplace"),
		field.Time("birthdate"),
		field.String("education"),
		field.String("job").Default("-"),
	}
}

// Edges of the Family.
func (Family) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

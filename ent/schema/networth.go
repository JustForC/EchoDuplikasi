package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Networth holds the schema definition for the Networth entity.
type Networth struct {
	ent.Schema
}

// Fields of the Networth.
func (Networth) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("value"),
	}
}

// Edges of the Networth.
func (Networth) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

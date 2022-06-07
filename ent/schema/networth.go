package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
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
	return []ent.Edge{
		edge.To("register", Register.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

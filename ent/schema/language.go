package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("talk"),
		field.String("write"),
		field.String("read"),
		field.String("listen"),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("register", Register.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

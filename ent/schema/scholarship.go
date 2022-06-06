package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Scholarship holds the schema definition for the Scholarship entity.
type Scholarship struct {
	ent.Schema
}

// Fields of the Scholarship.
func (Scholarship) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("startStepOne"),
		field.Time("startStepTwo"),
		field.Time("endStepOne"),
		field.Time("endStepTwo"),
		field.Time("announceStepOne"),
		field.Time("announceStepTwo"),
		field.String("onlineTest"),
		field.Int("status"),
	}
}

// Edges of the Scholarship.
func (Scholarship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("registers", Register.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Register holds the schema definition for the Register entity.
type Register struct {
	ent.Schema
}

// Fields of the Register.
func (Register) Fields() []ent.Field {
	return []ent.Field{
		field.Int("statusOne"),
		field.Int("statusTwo").Default(0),
		field.String("onlineInterview").Optional().Nillable(),
		field.Time("interviewTime").Optional().Nillable(),
	}
}

// Edges of the Register.
func (Register) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.To("scholarship", Scholarship.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
		edge.From("achievement", Achievement.Type).Ref("register"),
		edge.From("biodata", Biodata.Type).Ref("register"),
		edge.From("education", Education.Type).Ref("register"),
		edge.From("family", Family.Type).Ref("register"),
		edge.From("language", Language.Type).Ref("register"),
	}
}

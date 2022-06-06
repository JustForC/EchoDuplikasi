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
		// edge.From("user", User.Type).Ref("registers").Unique(),
		// edge.From("scholarship", Scholarship.Type).Ref("registers").Unique(),
		// edge.To("biodata", Biodata.Type).Unique().Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("networth", Networth.Type).Unique().Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("achievement", Achievement.Type).Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("education", Education.Type).Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("family", Family.Type).Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("language", Language.Type).Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("organization", Organization.Type).Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("socialmedia", SocialMedia.Type).Unique().Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("talent", Talent.Type).Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
		// edge.To("training", Training.Type).Annotations(entsql.Annotation{
		// 	OnDelete: entsql.Cascade,
		// }),
	}
}

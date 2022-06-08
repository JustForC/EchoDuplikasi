package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SocialMedia holds the schema definition for the SocialMedia entity.
type SocialMedia struct {
	ent.Schema
}

// Fields of the SocialMedia.
func (SocialMedia) Fields() []ent.Field {
	return []ent.Field{
		field.String("instagram"),
		field.String("facebook"),
		field.String("tiktok"),
		field.String("twitter"),
	}
}

// Edges of the SocialMedia.
func (SocialMedia) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("register", Register.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

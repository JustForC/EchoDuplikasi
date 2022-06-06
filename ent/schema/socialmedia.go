package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SocialMedia holds the schema definition for the SocialMedia entity.
type SocialMedia struct {
	ent.Schema
}

// Fields of the SocialMedia.
func (SocialMedia) Fields() []ent.Field {
	return []ent.Field{
		field.String("instagram").Optional().Nillable(),
		field.String("facebook").Optional().Nillable(),
		field.String("tiktok").Optional().Nillable(),
		field.String("twitter").Optional().Nillable(),
	}
}

// Edges of the SocialMedia.
func (SocialMedia) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{}
}

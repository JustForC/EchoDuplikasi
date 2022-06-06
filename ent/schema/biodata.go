package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Biodata holds the schema definition for the Biodata entity.
type Biodata struct {
	ent.Schema
}

// Fields of the Biodata.
func (Biodata) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("nickname"),
		field.String("gender"),
		field.String("birthplace"),
		field.Time("birthdate"),
		field.String("telephone"),
		field.String("email"),
		field.String("idType"),
		field.String("idNumber"),
		field.String("addressID"),
		field.String("postCodeID"),
		field.String("districtID"),
		field.String("cityID"),
		field.String("provinceID"),
		field.String("addressLiving"),
		field.String("postCodeLiving"),
		field.String("districtLiving"),
		field.String("cityLiving"),
		field.String("provinceLiving"),
		field.String("entrance"),
		field.String("entranceNumber"),
		field.String("major"),
		field.String("university"),
	}
}

// Edges of the Biodata.
func (Biodata) Edges() []ent.Edge {
	return nil
	// return []ent.Edge{
	// 	edge.From()
	// }
}

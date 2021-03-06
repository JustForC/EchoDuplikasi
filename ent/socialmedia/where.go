// Code generated by entc, DO NOT EDIT.

package socialmedia

import (
	"Kynesia/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Instagram applies equality check predicate on the "instagram" field. It's identical to InstagramEQ.
func Instagram(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInstagram), v))
	})
}

// Facebook applies equality check predicate on the "facebook" field. It's identical to FacebookEQ.
func Facebook(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFacebook), v))
	})
}

// Tiktok applies equality check predicate on the "tiktok" field. It's identical to TiktokEQ.
func Tiktok(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTiktok), v))
	})
}

// Twitter applies equality check predicate on the "twitter" field. It's identical to TwitterEQ.
func Twitter(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTwitter), v))
	})
}

// InstagramEQ applies the EQ predicate on the "instagram" field.
func InstagramEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInstagram), v))
	})
}

// InstagramNEQ applies the NEQ predicate on the "instagram" field.
func InstagramNEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInstagram), v))
	})
}

// InstagramIn applies the In predicate on the "instagram" field.
func InstagramIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInstagram), v...))
	})
}

// InstagramNotIn applies the NotIn predicate on the "instagram" field.
func InstagramNotIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInstagram), v...))
	})
}

// InstagramGT applies the GT predicate on the "instagram" field.
func InstagramGT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInstagram), v))
	})
}

// InstagramGTE applies the GTE predicate on the "instagram" field.
func InstagramGTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInstagram), v))
	})
}

// InstagramLT applies the LT predicate on the "instagram" field.
func InstagramLT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInstagram), v))
	})
}

// InstagramLTE applies the LTE predicate on the "instagram" field.
func InstagramLTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInstagram), v))
	})
}

// InstagramContains applies the Contains predicate on the "instagram" field.
func InstagramContains(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInstagram), v))
	})
}

// InstagramHasPrefix applies the HasPrefix predicate on the "instagram" field.
func InstagramHasPrefix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInstagram), v))
	})
}

// InstagramHasSuffix applies the HasSuffix predicate on the "instagram" field.
func InstagramHasSuffix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInstagram), v))
	})
}

// InstagramEqualFold applies the EqualFold predicate on the "instagram" field.
func InstagramEqualFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInstagram), v))
	})
}

// InstagramContainsFold applies the ContainsFold predicate on the "instagram" field.
func InstagramContainsFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInstagram), v))
	})
}

// FacebookEQ applies the EQ predicate on the "facebook" field.
func FacebookEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFacebook), v))
	})
}

// FacebookNEQ applies the NEQ predicate on the "facebook" field.
func FacebookNEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFacebook), v))
	})
}

// FacebookIn applies the In predicate on the "facebook" field.
func FacebookIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFacebook), v...))
	})
}

// FacebookNotIn applies the NotIn predicate on the "facebook" field.
func FacebookNotIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFacebook), v...))
	})
}

// FacebookGT applies the GT predicate on the "facebook" field.
func FacebookGT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFacebook), v))
	})
}

// FacebookGTE applies the GTE predicate on the "facebook" field.
func FacebookGTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFacebook), v))
	})
}

// FacebookLT applies the LT predicate on the "facebook" field.
func FacebookLT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFacebook), v))
	})
}

// FacebookLTE applies the LTE predicate on the "facebook" field.
func FacebookLTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFacebook), v))
	})
}

// FacebookContains applies the Contains predicate on the "facebook" field.
func FacebookContains(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFacebook), v))
	})
}

// FacebookHasPrefix applies the HasPrefix predicate on the "facebook" field.
func FacebookHasPrefix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFacebook), v))
	})
}

// FacebookHasSuffix applies the HasSuffix predicate on the "facebook" field.
func FacebookHasSuffix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFacebook), v))
	})
}

// FacebookEqualFold applies the EqualFold predicate on the "facebook" field.
func FacebookEqualFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFacebook), v))
	})
}

// FacebookContainsFold applies the ContainsFold predicate on the "facebook" field.
func FacebookContainsFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFacebook), v))
	})
}

// TiktokEQ applies the EQ predicate on the "tiktok" field.
func TiktokEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTiktok), v))
	})
}

// TiktokNEQ applies the NEQ predicate on the "tiktok" field.
func TiktokNEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTiktok), v))
	})
}

// TiktokIn applies the In predicate on the "tiktok" field.
func TiktokIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTiktok), v...))
	})
}

// TiktokNotIn applies the NotIn predicate on the "tiktok" field.
func TiktokNotIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTiktok), v...))
	})
}

// TiktokGT applies the GT predicate on the "tiktok" field.
func TiktokGT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTiktok), v))
	})
}

// TiktokGTE applies the GTE predicate on the "tiktok" field.
func TiktokGTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTiktok), v))
	})
}

// TiktokLT applies the LT predicate on the "tiktok" field.
func TiktokLT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTiktok), v))
	})
}

// TiktokLTE applies the LTE predicate on the "tiktok" field.
func TiktokLTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTiktok), v))
	})
}

// TiktokContains applies the Contains predicate on the "tiktok" field.
func TiktokContains(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTiktok), v))
	})
}

// TiktokHasPrefix applies the HasPrefix predicate on the "tiktok" field.
func TiktokHasPrefix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTiktok), v))
	})
}

// TiktokHasSuffix applies the HasSuffix predicate on the "tiktok" field.
func TiktokHasSuffix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTiktok), v))
	})
}

// TiktokEqualFold applies the EqualFold predicate on the "tiktok" field.
func TiktokEqualFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTiktok), v))
	})
}

// TiktokContainsFold applies the ContainsFold predicate on the "tiktok" field.
func TiktokContainsFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTiktok), v))
	})
}

// TwitterEQ applies the EQ predicate on the "twitter" field.
func TwitterEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTwitter), v))
	})
}

// TwitterNEQ applies the NEQ predicate on the "twitter" field.
func TwitterNEQ(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTwitter), v))
	})
}

// TwitterIn applies the In predicate on the "twitter" field.
func TwitterIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTwitter), v...))
	})
}

// TwitterNotIn applies the NotIn predicate on the "twitter" field.
func TwitterNotIn(vs ...string) predicate.SocialMedia {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SocialMedia(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTwitter), v...))
	})
}

// TwitterGT applies the GT predicate on the "twitter" field.
func TwitterGT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTwitter), v))
	})
}

// TwitterGTE applies the GTE predicate on the "twitter" field.
func TwitterGTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTwitter), v))
	})
}

// TwitterLT applies the LT predicate on the "twitter" field.
func TwitterLT(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTwitter), v))
	})
}

// TwitterLTE applies the LTE predicate on the "twitter" field.
func TwitterLTE(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTwitter), v))
	})
}

// TwitterContains applies the Contains predicate on the "twitter" field.
func TwitterContains(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTwitter), v))
	})
}

// TwitterHasPrefix applies the HasPrefix predicate on the "twitter" field.
func TwitterHasPrefix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTwitter), v))
	})
}

// TwitterHasSuffix applies the HasSuffix predicate on the "twitter" field.
func TwitterHasSuffix(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTwitter), v))
	})
}

// TwitterEqualFold applies the EqualFold predicate on the "twitter" field.
func TwitterEqualFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTwitter), v))
	})
}

// TwitterContainsFold applies the ContainsFold predicate on the "twitter" field.
func TwitterContainsFold(v string) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTwitter), v))
	})
}

// HasRegister applies the HasEdge predicate on the "register" edge.
func HasRegister() predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RegisterTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RegisterTable, RegisterPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegisterWith applies the HasEdge predicate on the "register" edge with a given conditions (other predicates).
func HasRegisterWith(preds ...predicate.Register) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RegisterInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RegisterTable, RegisterPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SocialMedia) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SocialMedia) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SocialMedia) predicate.SocialMedia {
	return predicate.SocialMedia(func(s *sql.Selector) {
		p(s.Not())
	})
}

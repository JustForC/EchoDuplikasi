// Code generated by entc, DO NOT EDIT.

package organization

import (
	"Kynesia/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Period applies equality check predicate on the "period" field. It's identical to PeriodEQ.
func Period(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriod), v))
	})
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PeriodEQ applies the EQ predicate on the "period" field.
func PeriodEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPeriod), v))
	})
}

// PeriodNEQ applies the NEQ predicate on the "period" field.
func PeriodNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPeriod), v))
	})
}

// PeriodIn applies the In predicate on the "period" field.
func PeriodIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPeriod), v...))
	})
}

// PeriodNotIn applies the NotIn predicate on the "period" field.
func PeriodNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPeriod), v...))
	})
}

// PeriodGT applies the GT predicate on the "period" field.
func PeriodGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPeriod), v))
	})
}

// PeriodGTE applies the GTE predicate on the "period" field.
func PeriodGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPeriod), v))
	})
}

// PeriodLT applies the LT predicate on the "period" field.
func PeriodLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPeriod), v))
	})
}

// PeriodLTE applies the LTE predicate on the "period" field.
func PeriodLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPeriod), v))
	})
}

// PeriodContains applies the Contains predicate on the "period" field.
func PeriodContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPeriod), v))
	})
}

// PeriodHasPrefix applies the HasPrefix predicate on the "period" field.
func PeriodHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPeriod), v))
	})
}

// PeriodHasSuffix applies the HasSuffix predicate on the "period" field.
func PeriodHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPeriod), v))
	})
}

// PeriodEqualFold applies the EqualFold predicate on the "period" field.
func PeriodEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPeriod), v))
	})
}

// PeriodContainsFold applies the ContainsFold predicate on the "period" field.
func PeriodContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPeriod), v))
	})
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPosition), v))
	})
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPosition), v...))
	})
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPosition), v...))
	})
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPosition), v))
	})
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPosition), v))
	})
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPosition), v))
	})
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPosition), v))
	})
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPosition), v))
	})
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPosition), v))
	})
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPosition), v))
	})
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPosition), v))
	})
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPosition), v))
	})
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetail), v))
	})
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetail), v...))
	})
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.Organization {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Organization(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetail), v...))
	})
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetail), v))
	})
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetail), v))
	})
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetail), v))
	})
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetail), v))
	})
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetail), v))
	})
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetail), v))
	})
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetail), v))
	})
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetail), v))
	})
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetail), v))
	})
}

// HasRegister applies the HasEdge predicate on the "register" edge.
func HasRegister() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RegisterTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RegisterTable, RegisterPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegisterWith applies the HasEdge predicate on the "register" edge with a given conditions (other predicates).
func HasRegisterWith(preds ...predicate.Register) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
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
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
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
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		p(s.Not())
	})
}

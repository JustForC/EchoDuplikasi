// Code generated by entc, DO NOT EDIT.

package register

import (
	"Kynesia/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StatusOne applies equality check predicate on the "statusOne" field. It's identical to StatusOneEQ.
func StatusOne(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusOne), v))
	})
}

// StatusTwo applies equality check predicate on the "statusTwo" field. It's identical to StatusTwoEQ.
func StatusTwo(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusTwo), v))
	})
}

// OnlineInterview applies equality check predicate on the "onlineInterview" field. It's identical to OnlineInterviewEQ.
func OnlineInterview(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnlineInterview), v))
	})
}

// InterviewTime applies equality check predicate on the "interviewTime" field. It's identical to InterviewTimeEQ.
func InterviewTime(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterviewTime), v))
	})
}

// StatusOneEQ applies the EQ predicate on the "statusOne" field.
func StatusOneEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusOne), v))
	})
}

// StatusOneNEQ applies the NEQ predicate on the "statusOne" field.
func StatusOneNEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusOne), v))
	})
}

// StatusOneIn applies the In predicate on the "statusOne" field.
func StatusOneIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusOne), v...))
	})
}

// StatusOneNotIn applies the NotIn predicate on the "statusOne" field.
func StatusOneNotIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusOne), v...))
	})
}

// StatusOneGT applies the GT predicate on the "statusOne" field.
func StatusOneGT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatusOne), v))
	})
}

// StatusOneGTE applies the GTE predicate on the "statusOne" field.
func StatusOneGTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatusOne), v))
	})
}

// StatusOneLT applies the LT predicate on the "statusOne" field.
func StatusOneLT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatusOne), v))
	})
}

// StatusOneLTE applies the LTE predicate on the "statusOne" field.
func StatusOneLTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatusOne), v))
	})
}

// StatusTwoEQ applies the EQ predicate on the "statusTwo" field.
func StatusTwoEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusTwo), v))
	})
}

// StatusTwoNEQ applies the NEQ predicate on the "statusTwo" field.
func StatusTwoNEQ(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusTwo), v))
	})
}

// StatusTwoIn applies the In predicate on the "statusTwo" field.
func StatusTwoIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusTwo), v...))
	})
}

// StatusTwoNotIn applies the NotIn predicate on the "statusTwo" field.
func StatusTwoNotIn(vs ...int) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusTwo), v...))
	})
}

// StatusTwoGT applies the GT predicate on the "statusTwo" field.
func StatusTwoGT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatusTwo), v))
	})
}

// StatusTwoGTE applies the GTE predicate on the "statusTwo" field.
func StatusTwoGTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatusTwo), v))
	})
}

// StatusTwoLT applies the LT predicate on the "statusTwo" field.
func StatusTwoLT(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatusTwo), v))
	})
}

// StatusTwoLTE applies the LTE predicate on the "statusTwo" field.
func StatusTwoLTE(v int) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatusTwo), v))
	})
}

// OnlineInterviewEQ applies the EQ predicate on the "onlineInterview" field.
func OnlineInterviewEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewNEQ applies the NEQ predicate on the "onlineInterview" field.
func OnlineInterviewNEQ(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewIn applies the In predicate on the "onlineInterview" field.
func OnlineInterviewIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOnlineInterview), v...))
	})
}

// OnlineInterviewNotIn applies the NotIn predicate on the "onlineInterview" field.
func OnlineInterviewNotIn(vs ...string) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOnlineInterview), v...))
	})
}

// OnlineInterviewGT applies the GT predicate on the "onlineInterview" field.
func OnlineInterviewGT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewGTE applies the GTE predicate on the "onlineInterview" field.
func OnlineInterviewGTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewLT applies the LT predicate on the "onlineInterview" field.
func OnlineInterviewLT(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewLTE applies the LTE predicate on the "onlineInterview" field.
func OnlineInterviewLTE(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewContains applies the Contains predicate on the "onlineInterview" field.
func OnlineInterviewContains(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewHasPrefix applies the HasPrefix predicate on the "onlineInterview" field.
func OnlineInterviewHasPrefix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewHasSuffix applies the HasSuffix predicate on the "onlineInterview" field.
func OnlineInterviewHasSuffix(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewIsNil applies the IsNil predicate on the "onlineInterview" field.
func OnlineInterviewIsNil() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOnlineInterview)))
	})
}

// OnlineInterviewNotNil applies the NotNil predicate on the "onlineInterview" field.
func OnlineInterviewNotNil() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOnlineInterview)))
	})
}

// OnlineInterviewEqualFold applies the EqualFold predicate on the "onlineInterview" field.
func OnlineInterviewEqualFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOnlineInterview), v))
	})
}

// OnlineInterviewContainsFold applies the ContainsFold predicate on the "onlineInterview" field.
func OnlineInterviewContainsFold(v string) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOnlineInterview), v))
	})
}

// InterviewTimeEQ applies the EQ predicate on the "interviewTime" field.
func InterviewTimeEQ(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInterviewTime), v))
	})
}

// InterviewTimeNEQ applies the NEQ predicate on the "interviewTime" field.
func InterviewTimeNEQ(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInterviewTime), v))
	})
}

// InterviewTimeIn applies the In predicate on the "interviewTime" field.
func InterviewTimeIn(vs ...time.Time) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInterviewTime), v...))
	})
}

// InterviewTimeNotIn applies the NotIn predicate on the "interviewTime" field.
func InterviewTimeNotIn(vs ...time.Time) predicate.Register {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Register(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInterviewTime), v...))
	})
}

// InterviewTimeGT applies the GT predicate on the "interviewTime" field.
func InterviewTimeGT(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInterviewTime), v))
	})
}

// InterviewTimeGTE applies the GTE predicate on the "interviewTime" field.
func InterviewTimeGTE(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInterviewTime), v))
	})
}

// InterviewTimeLT applies the LT predicate on the "interviewTime" field.
func InterviewTimeLT(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInterviewTime), v))
	})
}

// InterviewTimeLTE applies the LTE predicate on the "interviewTime" field.
func InterviewTimeLTE(v time.Time) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInterviewTime), v))
	})
}

// InterviewTimeIsNil applies the IsNil predicate on the "interviewTime" field.
func InterviewTimeIsNil() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInterviewTime)))
	})
}

// InterviewTimeNotNil applies the NotNil predicate on the "interviewTime" field.
func InterviewTimeNotNil() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInterviewTime)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasScholarship applies the HasEdge predicate on the "scholarship" edge.
func HasScholarship() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ScholarshipTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ScholarshipTable, ScholarshipPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScholarshipWith applies the HasEdge predicate on the "scholarship" edge with a given conditions (other predicates).
func HasScholarshipWith(preds ...predicate.Scholarship) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ScholarshipInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ScholarshipTable, ScholarshipPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAchievement applies the HasEdge predicate on the "achievement" edge.
func HasAchievement() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AchievementTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AchievementTable, AchievementPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAchievementWith applies the HasEdge predicate on the "achievement" edge with a given conditions (other predicates).
func HasAchievementWith(preds ...predicate.Achievement) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AchievementInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AchievementTable, AchievementPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBiodata applies the HasEdge predicate on the "biodata" edge.
func HasBiodata() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BiodataTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BiodataTable, BiodataPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBiodataWith applies the HasEdge predicate on the "biodata" edge with a given conditions (other predicates).
func HasBiodataWith(preds ...predicate.Biodata) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BiodataInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BiodataTable, BiodataPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEducation applies the HasEdge predicate on the "education" edge.
func HasEducation() predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EducationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, EducationTable, EducationPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEducationWith applies the HasEdge predicate on the "education" edge with a given conditions (other predicates).
func HasEducationWith(preds ...predicate.Education) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EducationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, EducationTable, EducationPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Register) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Register) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
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
func Not(p predicate.Register) predicate.Register {
	return predicate.Register(func(s *sql.Selector) {
		p(s.Not())
	})
}

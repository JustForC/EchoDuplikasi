// Code generated by entc, DO NOT EDIT.

package education

import (
	"Kynesia/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Grade applies equality check predicate on the "grade" field. It's identical to GradeEQ.
func Grade(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGrade), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Province applies equality check predicate on the "province" field. It's identical to ProvinceEQ.
func Province(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvince), v))
	})
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// Enter applies equality check predicate on the "enter" field. It's identical to EnterEQ.
func Enter(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnter), v))
	})
}

// Graduate applies equality check predicate on the "graduate" field. It's identical to GraduateEQ.
func Graduate(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGraduate), v))
	})
}

// GradeEQ applies the EQ predicate on the "grade" field.
func GradeEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGrade), v))
	})
}

// GradeNEQ applies the NEQ predicate on the "grade" field.
func GradeNEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGrade), v))
	})
}

// GradeIn applies the In predicate on the "grade" field.
func GradeIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGrade), v...))
	})
}

// GradeNotIn applies the NotIn predicate on the "grade" field.
func GradeNotIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGrade), v...))
	})
}

// GradeGT applies the GT predicate on the "grade" field.
func GradeGT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGrade), v))
	})
}

// GradeGTE applies the GTE predicate on the "grade" field.
func GradeGTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGrade), v))
	})
}

// GradeLT applies the LT predicate on the "grade" field.
func GradeLT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGrade), v))
	})
}

// GradeLTE applies the LTE predicate on the "grade" field.
func GradeLTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGrade), v))
	})
}

// GradeContains applies the Contains predicate on the "grade" field.
func GradeContains(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGrade), v))
	})
}

// GradeHasPrefix applies the HasPrefix predicate on the "grade" field.
func GradeHasPrefix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGrade), v))
	})
}

// GradeHasSuffix applies the HasSuffix predicate on the "grade" field.
func GradeHasSuffix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGrade), v))
	})
}

// GradeEqualFold applies the EqualFold predicate on the "grade" field.
func GradeEqualFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGrade), v))
	})
}

// GradeContainsFold applies the ContainsFold predicate on the "grade" field.
func GradeContainsFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGrade), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
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
func NameGT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ProvinceEQ applies the EQ predicate on the "province" field.
func ProvinceEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvince), v))
	})
}

// ProvinceNEQ applies the NEQ predicate on the "province" field.
func ProvinceNEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProvince), v))
	})
}

// ProvinceIn applies the In predicate on the "province" field.
func ProvinceIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProvince), v...))
	})
}

// ProvinceNotIn applies the NotIn predicate on the "province" field.
func ProvinceNotIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProvince), v...))
	})
}

// ProvinceGT applies the GT predicate on the "province" field.
func ProvinceGT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProvince), v))
	})
}

// ProvinceGTE applies the GTE predicate on the "province" field.
func ProvinceGTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProvince), v))
	})
}

// ProvinceLT applies the LT predicate on the "province" field.
func ProvinceLT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProvince), v))
	})
}

// ProvinceLTE applies the LTE predicate on the "province" field.
func ProvinceLTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProvince), v))
	})
}

// ProvinceContains applies the Contains predicate on the "province" field.
func ProvinceContains(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProvince), v))
	})
}

// ProvinceHasPrefix applies the HasPrefix predicate on the "province" field.
func ProvinceHasPrefix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProvince), v))
	})
}

// ProvinceHasSuffix applies the HasSuffix predicate on the "province" field.
func ProvinceHasSuffix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProvince), v))
	})
}

// ProvinceEqualFold applies the EqualFold predicate on the "province" field.
func ProvinceEqualFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProvince), v))
	})
}

// ProvinceContainsFold applies the ContainsFold predicate on the "province" field.
func ProvinceContainsFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProvince), v))
	})
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCity), v))
	})
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCity), v...))
	})
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCity), v...))
	})
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCity), v))
	})
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCity), v))
	})
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCity), v))
	})
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCity), v))
	})
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCity), v))
	})
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCity), v))
	})
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCity), v))
	})
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCity), v))
	})
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCity), v))
	})
}

// EnterEQ applies the EQ predicate on the "enter" field.
func EnterEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnter), v))
	})
}

// EnterNEQ applies the NEQ predicate on the "enter" field.
func EnterNEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnter), v))
	})
}

// EnterIn applies the In predicate on the "enter" field.
func EnterIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEnter), v...))
	})
}

// EnterNotIn applies the NotIn predicate on the "enter" field.
func EnterNotIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEnter), v...))
	})
}

// EnterGT applies the GT predicate on the "enter" field.
func EnterGT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnter), v))
	})
}

// EnterGTE applies the GTE predicate on the "enter" field.
func EnterGTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnter), v))
	})
}

// EnterLT applies the LT predicate on the "enter" field.
func EnterLT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnter), v))
	})
}

// EnterLTE applies the LTE predicate on the "enter" field.
func EnterLTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnter), v))
	})
}

// EnterContains applies the Contains predicate on the "enter" field.
func EnterContains(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEnter), v))
	})
}

// EnterHasPrefix applies the HasPrefix predicate on the "enter" field.
func EnterHasPrefix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEnter), v))
	})
}

// EnterHasSuffix applies the HasSuffix predicate on the "enter" field.
func EnterHasSuffix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEnter), v))
	})
}

// EnterEqualFold applies the EqualFold predicate on the "enter" field.
func EnterEqualFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEnter), v))
	})
}

// EnterContainsFold applies the ContainsFold predicate on the "enter" field.
func EnterContainsFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEnter), v))
	})
}

// GraduateEQ applies the EQ predicate on the "graduate" field.
func GraduateEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGraduate), v))
	})
}

// GraduateNEQ applies the NEQ predicate on the "graduate" field.
func GraduateNEQ(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGraduate), v))
	})
}

// GraduateIn applies the In predicate on the "graduate" field.
func GraduateIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGraduate), v...))
	})
}

// GraduateNotIn applies the NotIn predicate on the "graduate" field.
func GraduateNotIn(vs ...string) predicate.Education {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Education(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGraduate), v...))
	})
}

// GraduateGT applies the GT predicate on the "graduate" field.
func GraduateGT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGraduate), v))
	})
}

// GraduateGTE applies the GTE predicate on the "graduate" field.
func GraduateGTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGraduate), v))
	})
}

// GraduateLT applies the LT predicate on the "graduate" field.
func GraduateLT(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGraduate), v))
	})
}

// GraduateLTE applies the LTE predicate on the "graduate" field.
func GraduateLTE(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGraduate), v))
	})
}

// GraduateContains applies the Contains predicate on the "graduate" field.
func GraduateContains(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGraduate), v))
	})
}

// GraduateHasPrefix applies the HasPrefix predicate on the "graduate" field.
func GraduateHasPrefix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGraduate), v))
	})
}

// GraduateHasSuffix applies the HasSuffix predicate on the "graduate" field.
func GraduateHasSuffix(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGraduate), v))
	})
}

// GraduateEqualFold applies the EqualFold predicate on the "graduate" field.
func GraduateEqualFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGraduate), v))
	})
}

// GraduateContainsFold applies the ContainsFold predicate on the "graduate" field.
func GraduateContainsFold(v string) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGraduate), v))
	})
}

// HasRegister applies the HasEdge predicate on the "register" edge.
func HasRegister() predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RegisterTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RegisterTable, RegisterPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRegisterWith applies the HasEdge predicate on the "register" edge with a given conditions (other predicates).
func HasRegisterWith(preds ...predicate.Register) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
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
func And(predicates ...predicate.Education) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Education) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
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
func Not(p predicate.Education) predicate.Education {
	return predicate.Education(func(s *sql.Selector) {
		p(s.Not())
	})
}

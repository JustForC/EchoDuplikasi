// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/training"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Training is the model entity for the Training schema.
type Training struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Period holds the value of the "period" field.
	Period *string `json:"period,omitempty"`
	// Year holds the value of the "year" field.
	Year *string `json:"year,omitempty"`
	// Organizer holds the value of the "organizer" field.
	Organizer *string `json:"organizer,omitempty"`
	// Certificate holds the value of the "certificate" field.
	Certificate *string `json:"certificate,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Training) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case training.FieldID:
			values[i] = new(sql.NullInt64)
		case training.FieldName, training.FieldPeriod, training.FieldYear, training.FieldOrganizer, training.FieldCertificate:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Training", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Training fields.
func (t *Training) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case training.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case training.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = new(string)
				*t.Name = value.String
			}
		case training.FieldPeriod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field period", values[i])
			} else if value.Valid {
				t.Period = new(string)
				*t.Period = value.String
			}
		case training.FieldYear:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field year", values[i])
			} else if value.Valid {
				t.Year = new(string)
				*t.Year = value.String
			}
		case training.FieldOrganizer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organizer", values[i])
			} else if value.Valid {
				t.Organizer = new(string)
				*t.Organizer = value.String
			}
		case training.FieldCertificate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field certificate", values[i])
			} else if value.Valid {
				t.Certificate = new(string)
				*t.Certificate = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Training.
// Note that you need to call Training.Unwrap() before calling this method if this Training
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Training) Update() *TrainingUpdateOne {
	return (&TrainingClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Training entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Training) Unwrap() *Training {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Training is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Training) String() string {
	var builder strings.Builder
	builder.WriteString("Training(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	if v := t.Name; v != nil {
		builder.WriteString(", name=")
		builder.WriteString(*v)
	}
	if v := t.Period; v != nil {
		builder.WriteString(", period=")
		builder.WriteString(*v)
	}
	if v := t.Year; v != nil {
		builder.WriteString(", year=")
		builder.WriteString(*v)
	}
	if v := t.Organizer; v != nil {
		builder.WriteString(", organizer=")
		builder.WriteString(*v)
	}
	if v := t.Certificate; v != nil {
		builder.WriteString(", certificate=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Trainings is a parsable slice of Training.
type Trainings []*Training

func (t Trainings) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}

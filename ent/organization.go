// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/organization"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Period holds the value of the "period" field.
	Period *string `json:"period,omitempty"`
	// Position holds the value of the "position" field.
	Position *string `json:"position,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail *string `json:"detail,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			values[i] = new(sql.NullInt64)
		case organization.FieldName, organization.FieldPeriod, organization.FieldPosition, organization.FieldDetail:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Organization", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case organization.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = new(string)
				*o.Name = value.String
			}
		case organization.FieldPeriod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field period", values[i])
			} else if value.Valid {
				o.Period = new(string)
				*o.Period = value.String
			}
		case organization.FieldPosition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field position", values[i])
			} else if value.Valid {
				o.Position = new(string)
				*o.Position = value.String
			}
		case organization.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				o.Detail = new(string)
				*o.Detail = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return (&OrganizationClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	if v := o.Name; v != nil {
		builder.WriteString(", name=")
		builder.WriteString(*v)
	}
	if v := o.Period; v != nil {
		builder.WriteString(", period=")
		builder.WriteString(*v)
	}
	if v := o.Position; v != nil {
		builder.WriteString(", position=")
		builder.WriteString(*v)
	}
	if v := o.Detail; v != nil {
		builder.WriteString(", detail=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization

func (o Organizations) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/networth"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Networth is the model entity for the Networth schema.
type Networth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value int64 `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NetworthQuery when eager-loading is set.
	Edges NetworthEdges `json:"edges"`
}

// NetworthEdges holds the relations/edges for other nodes in the graph.
type NetworthEdges struct {
	// Register holds the value of the register edge.
	Register []*Register `json:"register,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RegisterOrErr returns the Register value or an error if the edge
// was not loaded in eager-loading.
func (e NetworthEdges) RegisterOrErr() ([]*Register, error) {
	if e.loadedTypes[0] {
		return e.Register, nil
	}
	return nil, &NotLoadedError{edge: "register"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Networth) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case networth.FieldID, networth.FieldValue:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Networth", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Networth fields.
func (n *Networth) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case networth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case networth.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				n.Value = value.Int64
			}
		}
	}
	return nil
}

// QueryRegister queries the "register" edge of the Networth entity.
func (n *Networth) QueryRegister() *RegisterQuery {
	return (&NetworthClient{config: n.config}).QueryRegister(n)
}

// Update returns a builder for updating this Networth.
// Note that you need to call Networth.Unwrap() before calling this method if this Networth
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Networth) Update() *NetworthUpdateOne {
	return (&NetworthClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the Networth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Networth) Unwrap() *Networth {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Networth is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Networth) String() string {
	var builder strings.Builder
	builder.WriteString("Networth(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", n.Value))
	builder.WriteByte(')')
	return builder.String()
}

// Networths is a parsable slice of Networth.
type Networths []*Networth

func (n Networths) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}

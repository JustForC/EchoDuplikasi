// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/socialmedia"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// SocialMedia is the model entity for the SocialMedia schema.
type SocialMedia struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Instagram holds the value of the "instagram" field.
	Instagram string `json:"instagram,omitempty"`
	// Facebook holds the value of the "facebook" field.
	Facebook string `json:"facebook,omitempty"`
	// Tiktok holds the value of the "tiktok" field.
	Tiktok string `json:"tiktok,omitempty"`
	// Twitter holds the value of the "twitter" field.
	Twitter string `json:"twitter,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SocialMediaQuery when eager-loading is set.
	Edges SocialMediaEdges `json:"edges"`
}

// SocialMediaEdges holds the relations/edges for other nodes in the graph.
type SocialMediaEdges struct {
	// Register holds the value of the register edge.
	Register []*Register `json:"register,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RegisterOrErr returns the Register value or an error if the edge
// was not loaded in eager-loading.
func (e SocialMediaEdges) RegisterOrErr() ([]*Register, error) {
	if e.loadedTypes[0] {
		return e.Register, nil
	}
	return nil, &NotLoadedError{edge: "register"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SocialMedia) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case socialmedia.FieldID:
			values[i] = new(sql.NullInt64)
		case socialmedia.FieldInstagram, socialmedia.FieldFacebook, socialmedia.FieldTiktok, socialmedia.FieldTwitter:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SocialMedia", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SocialMedia fields.
func (sm *SocialMedia) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case socialmedia.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = int(value.Int64)
		case socialmedia.FieldInstagram:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field instagram", values[i])
			} else if value.Valid {
				sm.Instagram = value.String
			}
		case socialmedia.FieldFacebook:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field facebook", values[i])
			} else if value.Valid {
				sm.Facebook = value.String
			}
		case socialmedia.FieldTiktok:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tiktok", values[i])
			} else if value.Valid {
				sm.Tiktok = value.String
			}
		case socialmedia.FieldTwitter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field twitter", values[i])
			} else if value.Valid {
				sm.Twitter = value.String
			}
		}
	}
	return nil
}

// QueryRegister queries the "register" edge of the SocialMedia entity.
func (sm *SocialMedia) QueryRegister() *RegisterQuery {
	return (&SocialMediaClient{config: sm.config}).QueryRegister(sm)
}

// Update returns a builder for updating this SocialMedia.
// Note that you need to call SocialMedia.Unwrap() before calling this method if this SocialMedia
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *SocialMedia) Update() *SocialMediaUpdateOne {
	return (&SocialMediaClient{config: sm.config}).UpdateOne(sm)
}

// Unwrap unwraps the SocialMedia entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *SocialMedia) Unwrap() *SocialMedia {
	tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: SocialMedia is not a transactional entity")
	}
	sm.config.driver = tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *SocialMedia) String() string {
	var builder strings.Builder
	builder.WriteString("SocialMedia(")
	builder.WriteString(fmt.Sprintf("id=%v", sm.ID))
	builder.WriteString(", instagram=")
	builder.WriteString(sm.Instagram)
	builder.WriteString(", facebook=")
	builder.WriteString(sm.Facebook)
	builder.WriteString(", tiktok=")
	builder.WriteString(sm.Tiktok)
	builder.WriteString(", twitter=")
	builder.WriteString(sm.Twitter)
	builder.WriteByte(')')
	return builder.String()
}

// SocialMediaSlice is a parsable slice of SocialMedia.
type SocialMediaSlice []*SocialMedia

func (sm SocialMediaSlice) config(cfg config) {
	for _i := range sm {
		sm[_i].config = cfg
	}
}

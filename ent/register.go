// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Kynesia/ent/register"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Register is the model entity for the Register schema.
type Register struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// StatusOne holds the value of the "statusOne" field.
	StatusOne int `json:"statusOne,omitempty"`
	// StatusTwo holds the value of the "statusTwo" field.
	StatusTwo int `json:"statusTwo,omitempty"`
	// OnlineInterview holds the value of the "onlineInterview" field.
	OnlineInterview *string `json:"onlineInterview,omitempty"`
	// InterviewTime holds the value of the "interviewTime" field.
	InterviewTime *time.Time `json:"interviewTime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RegisterQuery when eager-loading is set.
	Edges RegisterEdges `json:"edges"`
}

// RegisterEdges holds the relations/edges for other nodes in the graph.
type RegisterEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Scholarship holds the value of the scholarship edge.
	Scholarship []*Scholarship `json:"scholarship,omitempty"`
	// Achievement holds the value of the achievement edge.
	Achievement []*Achievement `json:"achievement,omitempty"`
	// Biodata holds the value of the biodata edge.
	Biodata []*Biodata `json:"biodata,omitempty"`
	// Education holds the value of the education edge.
	Education []*Education `json:"education,omitempty"`
	// Family holds the value of the family edge.
	Family []*Family `json:"family,omitempty"`
	// Language holds the value of the language edge.
	Language []*Language `json:"language,omitempty"`
	// Networth holds the value of the networth edge.
	Networth []*Networth `json:"networth,omitempty"`
	// Organization holds the value of the organization edge.
	Organization []*Organization `json:"organization,omitempty"`
	// Socialmedia holds the value of the socialmedia edge.
	Socialmedia []*SocialMedia `json:"socialmedia,omitempty"`
	// Talent holds the value of the talent edge.
	Talent []*Talent `json:"talent,omitempty"`
	// Training holds the value of the training edge.
	Training []*Training `json:"training,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [12]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ScholarshipOrErr returns the Scholarship value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) ScholarshipOrErr() ([]*Scholarship, error) {
	if e.loadedTypes[1] {
		return e.Scholarship, nil
	}
	return nil, &NotLoadedError{edge: "scholarship"}
}

// AchievementOrErr returns the Achievement value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) AchievementOrErr() ([]*Achievement, error) {
	if e.loadedTypes[2] {
		return e.Achievement, nil
	}
	return nil, &NotLoadedError{edge: "achievement"}
}

// BiodataOrErr returns the Biodata value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) BiodataOrErr() ([]*Biodata, error) {
	if e.loadedTypes[3] {
		return e.Biodata, nil
	}
	return nil, &NotLoadedError{edge: "biodata"}
}

// EducationOrErr returns the Education value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) EducationOrErr() ([]*Education, error) {
	if e.loadedTypes[4] {
		return e.Education, nil
	}
	return nil, &NotLoadedError{edge: "education"}
}

// FamilyOrErr returns the Family value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) FamilyOrErr() ([]*Family, error) {
	if e.loadedTypes[5] {
		return e.Family, nil
	}
	return nil, &NotLoadedError{edge: "family"}
}

// LanguageOrErr returns the Language value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) LanguageOrErr() ([]*Language, error) {
	if e.loadedTypes[6] {
		return e.Language, nil
	}
	return nil, &NotLoadedError{edge: "language"}
}

// NetworthOrErr returns the Networth value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) NetworthOrErr() ([]*Networth, error) {
	if e.loadedTypes[7] {
		return e.Networth, nil
	}
	return nil, &NotLoadedError{edge: "networth"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) OrganizationOrErr() ([]*Organization, error) {
	if e.loadedTypes[8] {
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// SocialmediaOrErr returns the Socialmedia value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) SocialmediaOrErr() ([]*SocialMedia, error) {
	if e.loadedTypes[9] {
		return e.Socialmedia, nil
	}
	return nil, &NotLoadedError{edge: "socialmedia"}
}

// TalentOrErr returns the Talent value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) TalentOrErr() ([]*Talent, error) {
	if e.loadedTypes[10] {
		return e.Talent, nil
	}
	return nil, &NotLoadedError{edge: "talent"}
}

// TrainingOrErr returns the Training value or an error if the edge
// was not loaded in eager-loading.
func (e RegisterEdges) TrainingOrErr() ([]*Training, error) {
	if e.loadedTypes[11] {
		return e.Training, nil
	}
	return nil, &NotLoadedError{edge: "training"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Register) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case register.FieldID, register.FieldStatusOne, register.FieldStatusTwo:
			values[i] = new(sql.NullInt64)
		case register.FieldOnlineInterview:
			values[i] = new(sql.NullString)
		case register.FieldInterviewTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Register", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Register fields.
func (r *Register) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case register.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case register.FieldStatusOne:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field statusOne", values[i])
			} else if value.Valid {
				r.StatusOne = int(value.Int64)
			}
		case register.FieldStatusTwo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field statusTwo", values[i])
			} else if value.Valid {
				r.StatusTwo = int(value.Int64)
			}
		case register.FieldOnlineInterview:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field onlineInterview", values[i])
			} else if value.Valid {
				r.OnlineInterview = new(string)
				*r.OnlineInterview = value.String
			}
		case register.FieldInterviewTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field interviewTime", values[i])
			} else if value.Valid {
				r.InterviewTime = new(time.Time)
				*r.InterviewTime = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Register entity.
func (r *Register) QueryUser() *UserQuery {
	return (&RegisterClient{config: r.config}).QueryUser(r)
}

// QueryScholarship queries the "scholarship" edge of the Register entity.
func (r *Register) QueryScholarship() *ScholarshipQuery {
	return (&RegisterClient{config: r.config}).QueryScholarship(r)
}

// QueryAchievement queries the "achievement" edge of the Register entity.
func (r *Register) QueryAchievement() *AchievementQuery {
	return (&RegisterClient{config: r.config}).QueryAchievement(r)
}

// QueryBiodata queries the "biodata" edge of the Register entity.
func (r *Register) QueryBiodata() *BiodataQuery {
	return (&RegisterClient{config: r.config}).QueryBiodata(r)
}

// QueryEducation queries the "education" edge of the Register entity.
func (r *Register) QueryEducation() *EducationQuery {
	return (&RegisterClient{config: r.config}).QueryEducation(r)
}

// QueryFamily queries the "family" edge of the Register entity.
func (r *Register) QueryFamily() *FamilyQuery {
	return (&RegisterClient{config: r.config}).QueryFamily(r)
}

// QueryLanguage queries the "language" edge of the Register entity.
func (r *Register) QueryLanguage() *LanguageQuery {
	return (&RegisterClient{config: r.config}).QueryLanguage(r)
}

// QueryNetworth queries the "networth" edge of the Register entity.
func (r *Register) QueryNetworth() *NetworthQuery {
	return (&RegisterClient{config: r.config}).QueryNetworth(r)
}

// QueryOrganization queries the "organization" edge of the Register entity.
func (r *Register) QueryOrganization() *OrganizationQuery {
	return (&RegisterClient{config: r.config}).QueryOrganization(r)
}

// QuerySocialmedia queries the "socialmedia" edge of the Register entity.
func (r *Register) QuerySocialmedia() *SocialMediaQuery {
	return (&RegisterClient{config: r.config}).QuerySocialmedia(r)
}

// QueryTalent queries the "talent" edge of the Register entity.
func (r *Register) QueryTalent() *TalentQuery {
	return (&RegisterClient{config: r.config}).QueryTalent(r)
}

// QueryTraining queries the "training" edge of the Register entity.
func (r *Register) QueryTraining() *TrainingQuery {
	return (&RegisterClient{config: r.config}).QueryTraining(r)
}

// Update returns a builder for updating this Register.
// Note that you need to call Register.Unwrap() before calling this method if this Register
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Register) Update() *RegisterUpdateOne {
	return (&RegisterClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Register entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Register) Unwrap() *Register {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Register is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Register) String() string {
	var builder strings.Builder
	builder.WriteString("Register(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", statusOne=")
	builder.WriteString(fmt.Sprintf("%v", r.StatusOne))
	builder.WriteString(", statusTwo=")
	builder.WriteString(fmt.Sprintf("%v", r.StatusTwo))
	if v := r.OnlineInterview; v != nil {
		builder.WriteString(", onlineInterview=")
		builder.WriteString(*v)
	}
	if v := r.InterviewTime; v != nil {
		builder.WriteString(", interviewTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Registers is a parsable slice of Register.
type Registers []*Register

func (r Registers) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}

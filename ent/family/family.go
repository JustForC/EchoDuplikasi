// Code generated by entc, DO NOT EDIT.

package family

const (
	// Label holds the string label denoting the family type in the database.
	Label = "family"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBirthplace holds the string denoting the birthplace field in the database.
	FieldBirthplace = "birthplace"
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldEducation holds the string denoting the education field in the database.
	FieldEducation = "education"
	// FieldJob holds the string denoting the job field in the database.
	FieldJob = "job"
	// EdgeRegister holds the string denoting the register edge name in mutations.
	EdgeRegister = "register"
	// Table holds the table name of the family in the database.
	Table = "families"
	// RegisterTable is the table that holds the register relation/edge. The primary key declared below.
	RegisterTable = "family_register"
	// RegisterInverseTable is the table name for the Register entity.
	// It exists in this package in order to avoid circular dependency with the "register" package.
	RegisterInverseTable = "registers"
)

// Columns holds all SQL columns for family fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldName,
	FieldGender,
	FieldBirthplace,
	FieldBirthdate,
	FieldEducation,
	FieldJob,
}

var (
	// RegisterPrimaryKey and RegisterColumn2 are the table columns denoting the
	// primary key for the register relation (M2M).
	RegisterPrimaryKey = []string{"family_id", "register_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

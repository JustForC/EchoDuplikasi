// Code generated by entc, DO NOT EDIT.

package talent

const (
	// Label holds the string label denoting the talent type in the database.
	Label = "talent"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeRegister holds the string denoting the register edge name in mutations.
	EdgeRegister = "register"
	// Table holds the table name of the talent in the database.
	Table = "talents"
	// RegisterTable is the table that holds the register relation/edge. The primary key declared below.
	RegisterTable = "talent_register"
	// RegisterInverseTable is the table name for the Register entity.
	// It exists in this package in order to avoid circular dependency with the "register" package.
	RegisterInverseTable = "registers"
)

// Columns holds all SQL columns for talent fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// RegisterPrimaryKey and RegisterColumn2 are the table columns denoting the
	// primary key for the register relation (M2M).
	RegisterPrimaryKey = []string{"talent_id", "register_id"}
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

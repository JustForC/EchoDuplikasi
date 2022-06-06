// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeRegisters holds the string denoting the registers edge name in mutations.
	EdgeRegisters = "registers"
	// Table holds the table name of the user in the database.
	Table = "users"
	// RegistersTable is the table that holds the registers relation/edge.
	RegistersTable = "registers"
	// RegistersInverseTable is the table name for the Register entity.
	// It exists in this package in order to avoid circular dependency with the "register" package.
	RegistersInverseTable = "registers"
	// RegistersColumn is the table column denoting the registers relation/edge.
	RegistersColumn = "user_registers"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
	FieldPassword,
	FieldRole,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

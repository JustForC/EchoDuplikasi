// Code generated by entc, DO NOT EDIT.

package talent

const (
	// Label holds the string label denoting the talent type in the database.
	Label = "talent"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the talent in the database.
	Table = "talents"
)

// Columns holds all SQL columns for talent fields.
var Columns = []string{
	FieldID,
	FieldName,
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

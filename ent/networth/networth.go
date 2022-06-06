// Code generated by entc, DO NOT EDIT.

package networth

const (
	// Label holds the string label denoting the networth type in the database.
	Label = "networth"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the networth in the database.
	Table = "networths"
)

// Columns holds all SQL columns for networth fields.
var Columns = []string{
	FieldID,
	FieldValue,
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
// Code generated by entc, DO NOT EDIT.

package biodata

const (
	// Label holds the string label denoting the biodata type in the database.
	Label = "biodata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the biodata in the database.
	Table = "biodata"
)

// Columns holds all SQL columns for biodata fields.
var Columns = []string{
	FieldID,
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

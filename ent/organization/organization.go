// Code generated by entc, DO NOT EDIT.

package organization

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPeriod holds the string denoting the period field in the database.
	FieldPeriod = "period"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPeriod,
	FieldPosition,
	FieldDetail,
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
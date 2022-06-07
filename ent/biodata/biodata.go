// Code generated by entc, DO NOT EDIT.

package biodata

const (
	// Label holds the string label denoting the biodata type in the database.
	Label = "biodata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBirthplace holds the string denoting the birthplace field in the database.
	FieldBirthplace = "birthplace"
	// FieldBirthdate holds the string denoting the birthdate field in the database.
	FieldBirthdate = "birthdate"
	// FieldTelephone holds the string denoting the telephone field in the database.
	FieldTelephone = "telephone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldIdType holds the string denoting the idtype field in the database.
	FieldIdType = "id_type"
	// FieldIdNumber holds the string denoting the idnumber field in the database.
	FieldIdNumber = "id_number"
	// FieldAddressID holds the string denoting the addressid field in the database.
	FieldAddressID = "address_id"
	// FieldPostCodeID holds the string denoting the postcodeid field in the database.
	FieldPostCodeID = "post_code_id"
	// FieldDistrictID holds the string denoting the districtid field in the database.
	FieldDistrictID = "district_id"
	// FieldCityID holds the string denoting the cityid field in the database.
	FieldCityID = "city_id"
	// FieldProvinceID holds the string denoting the provinceid field in the database.
	FieldProvinceID = "province_id"
	// FieldAddressLiving holds the string denoting the addressliving field in the database.
	FieldAddressLiving = "address_living"
	// FieldPostCodeLiving holds the string denoting the postcodeliving field in the database.
	FieldPostCodeLiving = "post_code_living"
	// FieldDistrictLiving holds the string denoting the districtliving field in the database.
	FieldDistrictLiving = "district_living"
	// FieldCityLiving holds the string denoting the cityliving field in the database.
	FieldCityLiving = "city_living"
	// FieldProvinceLiving holds the string denoting the provinceliving field in the database.
	FieldProvinceLiving = "province_living"
	// FieldEntrance holds the string denoting the entrance field in the database.
	FieldEntrance = "entrance"
	// FieldEntranceNumber holds the string denoting the entrancenumber field in the database.
	FieldEntranceNumber = "entrance_number"
	// FieldMajor holds the string denoting the major field in the database.
	FieldMajor = "major"
	// FieldUniversity holds the string denoting the university field in the database.
	FieldUniversity = "university"
	// EdgeRegister holds the string denoting the register edge name in mutations.
	EdgeRegister = "register"
	// Table holds the table name of the biodata in the database.
	Table = "biodata"
	// RegisterTable is the table that holds the register relation/edge. The primary key declared below.
	RegisterTable = "biodata_register"
	// RegisterInverseTable is the table name for the Register entity.
	// It exists in this package in order to avoid circular dependency with the "register" package.
	RegisterInverseTable = "registers"
)

// Columns holds all SQL columns for biodata fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNickname,
	FieldGender,
	FieldBirthplace,
	FieldBirthdate,
	FieldTelephone,
	FieldEmail,
	FieldIdType,
	FieldIdNumber,
	FieldAddressID,
	FieldPostCodeID,
	FieldDistrictID,
	FieldCityID,
	FieldProvinceID,
	FieldAddressLiving,
	FieldPostCodeLiving,
	FieldDistrictLiving,
	FieldCityLiving,
	FieldProvinceLiving,
	FieldEntrance,
	FieldEntranceNumber,
	FieldMajor,
	FieldUniversity,
}

var (
	// RegisterPrimaryKey and RegisterColumn2 are the table columns denoting the
	// primary key for the register relation (M2M).
	RegisterPrimaryKey = []string{"biodata_id", "register_id"}
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

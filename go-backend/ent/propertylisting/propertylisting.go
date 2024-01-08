// Code generated by ent, DO NOT EDIT.

package propertylisting

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the propertylisting type in the database.
	Label = "property_listing"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHouseType holds the string denoting the housetype field in the database.
	FieldHouseType = "house_type"
	// FieldArea holds the string denoting the area field in the database.
	FieldArea = "area"
	// FieldNumOfBedroom holds the string denoting the num_of_bedroom field in the database.
	FieldNumOfBedroom = "num_of_bedroom"
	// FieldLeaseDate holds the string denoting the lease_date field in the database.
	FieldLeaseDate = "lease_date"
	// FieldNumberOfData holds the string denoting the number_of_data field in the database.
	FieldNumberOfData = "number_of_data"
	// FieldRentalPrice holds the string denoting the rental_price field in the database.
	FieldRentalPrice = "rental_price"
	// FieldY holds the string denoting the y field in the database.
	FieldY = "y"
	// FieldX holds the string denoting the x field in the database.
	FieldX = "x"
	// FieldStreet holds the string denoting the street field in the database.
	FieldStreet = "street"
	// FieldDistrict holds the string denoting the district field in the database.
	FieldDistrict = "district"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the propertylisting in the database.
	Table = "property_listings"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_propertylistings"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for propertylisting fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldHouseType,
	FieldArea,
	FieldNumOfBedroom,
	FieldLeaseDate,
	FieldNumberOfData,
	FieldRentalPrice,
	FieldY,
	FieldX,
	FieldStreet,
	FieldDistrict,
	FieldUUID,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "property_listing_id"}
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

var (
	// DefaultArea holds the default value on creation for the "area" field.
	DefaultArea int
	// LeaseDateValidator is a validator for the "lease_date" field. It is called by the builders before save.
	LeaseDateValidator func(string) error
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
)

// OrderOption defines the ordering options for the PropertyListing queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByHouseType orders the results by the houseType field.
func ByHouseType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHouseType, opts...).ToFunc()
}

// ByArea orders the results by the area field.
func ByArea(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArea, opts...).ToFunc()
}

// ByNumOfBedroom orders the results by the num_of_bedroom field.
func ByNumOfBedroom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumOfBedroom, opts...).ToFunc()
}

// ByLeaseDate orders the results by the lease_date field.
func ByLeaseDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLeaseDate, opts...).ToFunc()
}

// ByNumberOfData orders the results by the number_of_data field.
func ByNumberOfData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberOfData, opts...).ToFunc()
}

// ByRentalPrice orders the results by the rental_price field.
func ByRentalPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRentalPrice, opts...).ToFunc()
}

// ByY orders the results by the y field.
func ByY(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldY, opts...).ToFunc()
}

// ByX orders the results by the x field.
func ByX(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldX, opts...).ToFunc()
}

// ByStreet orders the results by the street field.
func ByStreet(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStreet, opts...).ToFunc()
}

// ByDistrict orders the results by the district field.
func ByDistrict(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDistrict, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
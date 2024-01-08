// Code generated by ent, DO NOT EDIT.

package propertylisting

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NattapolChan/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldName, v))
}

// HouseType applies equality check predicate on the "houseType" field. It's identical to HouseTypeEQ.
func HouseType(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldHouseType, v))
}

// Area applies equality check predicate on the "area" field. It's identical to AreaEQ.
func Area(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldArea, v))
}

// NumOfBedroom applies equality check predicate on the "num_of_bedroom" field. It's identical to NumOfBedroomEQ.
func NumOfBedroom(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldNumOfBedroom, v))
}

// LeaseDate applies equality check predicate on the "lease_date" field. It's identical to LeaseDateEQ.
func LeaseDate(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldLeaseDate, v))
}

// NumberOfData applies equality check predicate on the "number_of_data" field. It's identical to NumberOfDataEQ.
func NumberOfData(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldNumberOfData, v))
}

// RentalPrice applies equality check predicate on the "rental_price" field. It's identical to RentalPriceEQ.
func RentalPrice(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldRentalPrice, v))
}

// Y applies equality check predicate on the "y" field. It's identical to YEQ.
func Y(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldY, v))
}

// X applies equality check predicate on the "x" field. It's identical to XEQ.
func X(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldX, v))
}

// Street applies equality check predicate on the "street" field. It's identical to StreetEQ.
func Street(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldStreet, v))
}

// District applies equality check predicate on the "district" field. It's identical to DistrictEQ.
func District(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldDistrict, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldUUID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContainsFold(FieldName, v))
}

// HouseTypeEQ applies the EQ predicate on the "houseType" field.
func HouseTypeEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldHouseType, v))
}

// HouseTypeNEQ applies the NEQ predicate on the "houseType" field.
func HouseTypeNEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldHouseType, v))
}

// HouseTypeIn applies the In predicate on the "houseType" field.
func HouseTypeIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldHouseType, vs...))
}

// HouseTypeNotIn applies the NotIn predicate on the "houseType" field.
func HouseTypeNotIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldHouseType, vs...))
}

// HouseTypeGT applies the GT predicate on the "houseType" field.
func HouseTypeGT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldHouseType, v))
}

// HouseTypeGTE applies the GTE predicate on the "houseType" field.
func HouseTypeGTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldHouseType, v))
}

// HouseTypeLT applies the LT predicate on the "houseType" field.
func HouseTypeLT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldHouseType, v))
}

// HouseTypeLTE applies the LTE predicate on the "houseType" field.
func HouseTypeLTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldHouseType, v))
}

// HouseTypeContains applies the Contains predicate on the "houseType" field.
func HouseTypeContains(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContains(FieldHouseType, v))
}

// HouseTypeHasPrefix applies the HasPrefix predicate on the "houseType" field.
func HouseTypeHasPrefix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasPrefix(FieldHouseType, v))
}

// HouseTypeHasSuffix applies the HasSuffix predicate on the "houseType" field.
func HouseTypeHasSuffix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasSuffix(FieldHouseType, v))
}

// HouseTypeEqualFold applies the EqualFold predicate on the "houseType" field.
func HouseTypeEqualFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEqualFold(FieldHouseType, v))
}

// HouseTypeContainsFold applies the ContainsFold predicate on the "houseType" field.
func HouseTypeContainsFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContainsFold(FieldHouseType, v))
}

// AreaEQ applies the EQ predicate on the "area" field.
func AreaEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldArea, v))
}

// AreaNEQ applies the NEQ predicate on the "area" field.
func AreaNEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldArea, v))
}

// AreaIn applies the In predicate on the "area" field.
func AreaIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldArea, vs...))
}

// AreaNotIn applies the NotIn predicate on the "area" field.
func AreaNotIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldArea, vs...))
}

// AreaGT applies the GT predicate on the "area" field.
func AreaGT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldArea, v))
}

// AreaGTE applies the GTE predicate on the "area" field.
func AreaGTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldArea, v))
}

// AreaLT applies the LT predicate on the "area" field.
func AreaLT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldArea, v))
}

// AreaLTE applies the LTE predicate on the "area" field.
func AreaLTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldArea, v))
}

// NumOfBedroomEQ applies the EQ predicate on the "num_of_bedroom" field.
func NumOfBedroomEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldNumOfBedroom, v))
}

// NumOfBedroomNEQ applies the NEQ predicate on the "num_of_bedroom" field.
func NumOfBedroomNEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldNumOfBedroom, v))
}

// NumOfBedroomIn applies the In predicate on the "num_of_bedroom" field.
func NumOfBedroomIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldNumOfBedroom, vs...))
}

// NumOfBedroomNotIn applies the NotIn predicate on the "num_of_bedroom" field.
func NumOfBedroomNotIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldNumOfBedroom, vs...))
}

// NumOfBedroomGT applies the GT predicate on the "num_of_bedroom" field.
func NumOfBedroomGT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldNumOfBedroom, v))
}

// NumOfBedroomGTE applies the GTE predicate on the "num_of_bedroom" field.
func NumOfBedroomGTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldNumOfBedroom, v))
}

// NumOfBedroomLT applies the LT predicate on the "num_of_bedroom" field.
func NumOfBedroomLT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldNumOfBedroom, v))
}

// NumOfBedroomLTE applies the LTE predicate on the "num_of_bedroom" field.
func NumOfBedroomLTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldNumOfBedroom, v))
}

// LeaseDateEQ applies the EQ predicate on the "lease_date" field.
func LeaseDateEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldLeaseDate, v))
}

// LeaseDateNEQ applies the NEQ predicate on the "lease_date" field.
func LeaseDateNEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldLeaseDate, v))
}

// LeaseDateIn applies the In predicate on the "lease_date" field.
func LeaseDateIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldLeaseDate, vs...))
}

// LeaseDateNotIn applies the NotIn predicate on the "lease_date" field.
func LeaseDateNotIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldLeaseDate, vs...))
}

// LeaseDateGT applies the GT predicate on the "lease_date" field.
func LeaseDateGT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldLeaseDate, v))
}

// LeaseDateGTE applies the GTE predicate on the "lease_date" field.
func LeaseDateGTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldLeaseDate, v))
}

// LeaseDateLT applies the LT predicate on the "lease_date" field.
func LeaseDateLT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldLeaseDate, v))
}

// LeaseDateLTE applies the LTE predicate on the "lease_date" field.
func LeaseDateLTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldLeaseDate, v))
}

// LeaseDateContains applies the Contains predicate on the "lease_date" field.
func LeaseDateContains(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContains(FieldLeaseDate, v))
}

// LeaseDateHasPrefix applies the HasPrefix predicate on the "lease_date" field.
func LeaseDateHasPrefix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasPrefix(FieldLeaseDate, v))
}

// LeaseDateHasSuffix applies the HasSuffix predicate on the "lease_date" field.
func LeaseDateHasSuffix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasSuffix(FieldLeaseDate, v))
}

// LeaseDateEqualFold applies the EqualFold predicate on the "lease_date" field.
func LeaseDateEqualFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEqualFold(FieldLeaseDate, v))
}

// LeaseDateContainsFold applies the ContainsFold predicate on the "lease_date" field.
func LeaseDateContainsFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContainsFold(FieldLeaseDate, v))
}

// NumberOfDataEQ applies the EQ predicate on the "number_of_data" field.
func NumberOfDataEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldNumberOfData, v))
}

// NumberOfDataNEQ applies the NEQ predicate on the "number_of_data" field.
func NumberOfDataNEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldNumberOfData, v))
}

// NumberOfDataIn applies the In predicate on the "number_of_data" field.
func NumberOfDataIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldNumberOfData, vs...))
}

// NumberOfDataNotIn applies the NotIn predicate on the "number_of_data" field.
func NumberOfDataNotIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldNumberOfData, vs...))
}

// NumberOfDataGT applies the GT predicate on the "number_of_data" field.
func NumberOfDataGT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldNumberOfData, v))
}

// NumberOfDataGTE applies the GTE predicate on the "number_of_data" field.
func NumberOfDataGTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldNumberOfData, v))
}

// NumberOfDataLT applies the LT predicate on the "number_of_data" field.
func NumberOfDataLT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldNumberOfData, v))
}

// NumberOfDataLTE applies the LTE predicate on the "number_of_data" field.
func NumberOfDataLTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldNumberOfData, v))
}

// RentalPriceEQ applies the EQ predicate on the "rental_price" field.
func RentalPriceEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldRentalPrice, v))
}

// RentalPriceNEQ applies the NEQ predicate on the "rental_price" field.
func RentalPriceNEQ(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldRentalPrice, v))
}

// RentalPriceIn applies the In predicate on the "rental_price" field.
func RentalPriceIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldRentalPrice, vs...))
}

// RentalPriceNotIn applies the NotIn predicate on the "rental_price" field.
func RentalPriceNotIn(vs ...int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldRentalPrice, vs...))
}

// RentalPriceGT applies the GT predicate on the "rental_price" field.
func RentalPriceGT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldRentalPrice, v))
}

// RentalPriceGTE applies the GTE predicate on the "rental_price" field.
func RentalPriceGTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldRentalPrice, v))
}

// RentalPriceLT applies the LT predicate on the "rental_price" field.
func RentalPriceLT(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldRentalPrice, v))
}

// RentalPriceLTE applies the LTE predicate on the "rental_price" field.
func RentalPriceLTE(v int) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldRentalPrice, v))
}

// YEQ applies the EQ predicate on the "y" field.
func YEQ(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldY, v))
}

// YNEQ applies the NEQ predicate on the "y" field.
func YNEQ(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldY, v))
}

// YIn applies the In predicate on the "y" field.
func YIn(vs ...float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldY, vs...))
}

// YNotIn applies the NotIn predicate on the "y" field.
func YNotIn(vs ...float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldY, vs...))
}

// YGT applies the GT predicate on the "y" field.
func YGT(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldY, v))
}

// YGTE applies the GTE predicate on the "y" field.
func YGTE(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldY, v))
}

// YLT applies the LT predicate on the "y" field.
func YLT(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldY, v))
}

// YLTE applies the LTE predicate on the "y" field.
func YLTE(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldY, v))
}

// XEQ applies the EQ predicate on the "x" field.
func XEQ(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldX, v))
}

// XNEQ applies the NEQ predicate on the "x" field.
func XNEQ(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldX, v))
}

// XIn applies the In predicate on the "x" field.
func XIn(vs ...float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldX, vs...))
}

// XNotIn applies the NotIn predicate on the "x" field.
func XNotIn(vs ...float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldX, vs...))
}

// XGT applies the GT predicate on the "x" field.
func XGT(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldX, v))
}

// XGTE applies the GTE predicate on the "x" field.
func XGTE(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldX, v))
}

// XLT applies the LT predicate on the "x" field.
func XLT(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldX, v))
}

// XLTE applies the LTE predicate on the "x" field.
func XLTE(v float64) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldX, v))
}

// StreetEQ applies the EQ predicate on the "street" field.
func StreetEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldStreet, v))
}

// StreetNEQ applies the NEQ predicate on the "street" field.
func StreetNEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldStreet, v))
}

// StreetIn applies the In predicate on the "street" field.
func StreetIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldStreet, vs...))
}

// StreetNotIn applies the NotIn predicate on the "street" field.
func StreetNotIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldStreet, vs...))
}

// StreetGT applies the GT predicate on the "street" field.
func StreetGT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldStreet, v))
}

// StreetGTE applies the GTE predicate on the "street" field.
func StreetGTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldStreet, v))
}

// StreetLT applies the LT predicate on the "street" field.
func StreetLT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldStreet, v))
}

// StreetLTE applies the LTE predicate on the "street" field.
func StreetLTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldStreet, v))
}

// StreetContains applies the Contains predicate on the "street" field.
func StreetContains(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContains(FieldStreet, v))
}

// StreetHasPrefix applies the HasPrefix predicate on the "street" field.
func StreetHasPrefix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasPrefix(FieldStreet, v))
}

// StreetHasSuffix applies the HasSuffix predicate on the "street" field.
func StreetHasSuffix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasSuffix(FieldStreet, v))
}

// StreetIsNil applies the IsNil predicate on the "street" field.
func StreetIsNil() predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIsNull(FieldStreet))
}

// StreetNotNil applies the NotNil predicate on the "street" field.
func StreetNotNil() predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotNull(FieldStreet))
}

// StreetEqualFold applies the EqualFold predicate on the "street" field.
func StreetEqualFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEqualFold(FieldStreet, v))
}

// StreetContainsFold applies the ContainsFold predicate on the "street" field.
func StreetContainsFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContainsFold(FieldStreet, v))
}

// DistrictEQ applies the EQ predicate on the "district" field.
func DistrictEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldDistrict, v))
}

// DistrictNEQ applies the NEQ predicate on the "district" field.
func DistrictNEQ(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldDistrict, v))
}

// DistrictIn applies the In predicate on the "district" field.
func DistrictIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldDistrict, vs...))
}

// DistrictNotIn applies the NotIn predicate on the "district" field.
func DistrictNotIn(vs ...string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldDistrict, vs...))
}

// DistrictGT applies the GT predicate on the "district" field.
func DistrictGT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldDistrict, v))
}

// DistrictGTE applies the GTE predicate on the "district" field.
func DistrictGTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldDistrict, v))
}

// DistrictLT applies the LT predicate on the "district" field.
func DistrictLT(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldDistrict, v))
}

// DistrictLTE applies the LTE predicate on the "district" field.
func DistrictLTE(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldDistrict, v))
}

// DistrictContains applies the Contains predicate on the "district" field.
func DistrictContains(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContains(FieldDistrict, v))
}

// DistrictHasPrefix applies the HasPrefix predicate on the "district" field.
func DistrictHasPrefix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasPrefix(FieldDistrict, v))
}

// DistrictHasSuffix applies the HasSuffix predicate on the "district" field.
func DistrictHasSuffix(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldHasSuffix(FieldDistrict, v))
}

// DistrictEqualFold applies the EqualFold predicate on the "district" field.
func DistrictEqualFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEqualFold(FieldDistrict, v))
}

// DistrictContainsFold applies the ContainsFold predicate on the "district" field.
func DistrictContainsFold(v string) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldContainsFold(FieldDistrict, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.PropertyListing {
	return predicate.PropertyListing(sql.FieldLTE(FieldUUID, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.PropertyListing {
	return predicate.PropertyListing(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.PropertyListing {
	return predicate.PropertyListing(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PropertyListing) predicate.PropertyListing {
	return predicate.PropertyListing(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PropertyListing) predicate.PropertyListing {
	return predicate.PropertyListing(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PropertyListing) predicate.PropertyListing {
	return predicate.PropertyListing(sql.NotPredicates(p))
}
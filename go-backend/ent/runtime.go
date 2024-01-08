// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NattapolChan/ent/propertylisting"
	"github.com/NattapolChan/ent/schema"
	"github.com/NattapolChan/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	propertylistingFields := schema.PropertyListing{}.Fields()
	_ = propertylistingFields
	// propertylistingDescArea is the schema descriptor for area field.
	propertylistingDescArea := propertylistingFields[2].Descriptor()
	// propertylisting.DefaultArea holds the default value on creation for the area field.
	propertylisting.DefaultArea = propertylistingDescArea.Default.(int)
	// propertylistingDescLeaseDate is the schema descriptor for lease_date field.
	propertylistingDescLeaseDate := propertylistingFields[4].Descriptor()
	// propertylisting.LeaseDateValidator is a validator for the "lease_date" field. It is called by the builders before save.
	propertylisting.LeaseDateValidator = propertylistingDescLeaseDate.Validators[0].(func(string) error)
	// propertylistingDescUUID is the schema descriptor for uuid field.
	propertylistingDescUUID := propertylistingFields[11].Descriptor()
	// propertylisting.DefaultUUID holds the default value on creation for the uuid field.
	propertylisting.DefaultUUID = propertylistingDescUUID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmailAddr is the schema descriptor for emailAddr field.
	userDescEmailAddr := userFields[2].Descriptor()
	// user.EmailAddrValidator is a validator for the "emailAddr" field. It is called by the builders before save.
	user.EmailAddrValidator = userDescEmailAddr.Validators[0].(func(string) error)
	// userDescFavorites is the schema descriptor for favorites field.
	userDescFavorites := userFields[3].Descriptor()
	// user.DefaultFavorites holds the default value on creation for the favorites field.
	user.DefaultFavorites = userDescFavorites.Default.([]string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
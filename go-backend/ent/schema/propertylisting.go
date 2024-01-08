package schema

import (
	"errors"
	"strconv"
	"entgo.io/ent"
	"github.com/google/uuid"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PropertyListing holds the schema definition for the PropertyListing entity.
type PropertyListing struct {
	ent.Schema
}

// Fields of the PropertyListing.
func (PropertyListing) Fields() []ent.Field {
	return []ent.Field{
		// name of the project
		field.String("name"),

		// property info
		field.String("houseType"),
		field.Int("area").
			Default(0),
		field.Int("num_of_bedroom"),
		field.String("lease_date"). // 0224 = February of 2024
			Validate(func (s string) error {
				intVar, err := strconv.Atoi(s)
				if err != nil {
					return errors.New("lease date must be a string of number")
				}
				if  intVar > 1300 { 
					return errors.New("lease date must have the first 2 digits <= 12") 
				}
				return nil
			}),
		field.Int("number_of_data"),
		field.Int("rental_price"), // only latest are shown
		
		// location info
		field.Float("y"),
		field.Float("x"),
		field.String("street").
			Optional(),
		field.String("district"),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
	}
}

// Edges of the PropertyListing.
func (PropertyListing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("propertylistings"),
	}
}


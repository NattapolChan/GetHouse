package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.String("emailAddr").NotEmpty(),
		field.JSON("favorites", []string{}).Default([]string{}),
    	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("propertylistings", PropertyListing.Type),
	}
}

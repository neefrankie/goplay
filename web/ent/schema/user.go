package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// An edge from User to Car deining that a user can
		// have 1 or more cars, but a car has only one owner.
		edge.To("cars", Car.Type),
		// Create an inverse-edge called "groups" of type `Group`
		// can reference it to the "users" edge (in Group schema)
		// explicitly using the `Ref` method
		edge.From("groups", Group.Type).
			Ref("users"),
	}
}

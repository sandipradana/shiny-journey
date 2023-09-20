package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Admin struct {
	ent.Schema
}

func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("email").
			Unique(),
		field.String("password").
			Sensitive(),
	}
}

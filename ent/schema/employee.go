package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Employee struct {
	ent.Schema
}

func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").
			Unique(),
	}
}

func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attendances", Attendance.Type),
	}
}

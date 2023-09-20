package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Attendance struct {
	ent.Schema
}

func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("attendance_date").SchemaType(map[string]string{
			dialect.MySQL: "date",
		}),
		field.Time("check_in_time").Optional().Nillable(),
		field.Time("check_out_time").Optional().Nillable(),
		field.Enum("status").Values("absent", "present"),
	}
}

func (Attendance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).
			Ref("attendances").
			Unique(),
	}
}

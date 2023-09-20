// Code generated by ent, DO NOT EDIT.

package attendance

import (
	"shiny-journey/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldID, id))
}

// AttendanceDate applies equality check predicate on the "attendance_date" field. It's identical to AttendanceDateEQ.
func AttendanceDate(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldAttendanceDate, v))
}

// CheckInTime applies equality check predicate on the "check_in_time" field. It's identical to CheckInTimeEQ.
func CheckInTime(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckInTime, v))
}

// CheckOutTime applies equality check predicate on the "check_out_time" field. It's identical to CheckOutTimeEQ.
func CheckOutTime(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckOutTime, v))
}

// AttendanceDateEQ applies the EQ predicate on the "attendance_date" field.
func AttendanceDateEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldAttendanceDate, v))
}

// AttendanceDateNEQ applies the NEQ predicate on the "attendance_date" field.
func AttendanceDateNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldAttendanceDate, v))
}

// AttendanceDateIn applies the In predicate on the "attendance_date" field.
func AttendanceDateIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldAttendanceDate, vs...))
}

// AttendanceDateNotIn applies the NotIn predicate on the "attendance_date" field.
func AttendanceDateNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldAttendanceDate, vs...))
}

// AttendanceDateGT applies the GT predicate on the "attendance_date" field.
func AttendanceDateGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldAttendanceDate, v))
}

// AttendanceDateGTE applies the GTE predicate on the "attendance_date" field.
func AttendanceDateGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldAttendanceDate, v))
}

// AttendanceDateLT applies the LT predicate on the "attendance_date" field.
func AttendanceDateLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldAttendanceDate, v))
}

// AttendanceDateLTE applies the LTE predicate on the "attendance_date" field.
func AttendanceDateLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldAttendanceDate, v))
}

// CheckInTimeEQ applies the EQ predicate on the "check_in_time" field.
func CheckInTimeEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckInTime, v))
}

// CheckInTimeNEQ applies the NEQ predicate on the "check_in_time" field.
func CheckInTimeNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldCheckInTime, v))
}

// CheckInTimeIn applies the In predicate on the "check_in_time" field.
func CheckInTimeIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldCheckInTime, vs...))
}

// CheckInTimeNotIn applies the NotIn predicate on the "check_in_time" field.
func CheckInTimeNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldCheckInTime, vs...))
}

// CheckInTimeGT applies the GT predicate on the "check_in_time" field.
func CheckInTimeGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldCheckInTime, v))
}

// CheckInTimeGTE applies the GTE predicate on the "check_in_time" field.
func CheckInTimeGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldCheckInTime, v))
}

// CheckInTimeLT applies the LT predicate on the "check_in_time" field.
func CheckInTimeLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldCheckInTime, v))
}

// CheckInTimeLTE applies the LTE predicate on the "check_in_time" field.
func CheckInTimeLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldCheckInTime, v))
}

// CheckInTimeIsNil applies the IsNil predicate on the "check_in_time" field.
func CheckInTimeIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldCheckInTime))
}

// CheckInTimeNotNil applies the NotNil predicate on the "check_in_time" field.
func CheckInTimeNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldCheckInTime))
}

// CheckOutTimeEQ applies the EQ predicate on the "check_out_time" field.
func CheckOutTimeEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldCheckOutTime, v))
}

// CheckOutTimeNEQ applies the NEQ predicate on the "check_out_time" field.
func CheckOutTimeNEQ(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldCheckOutTime, v))
}

// CheckOutTimeIn applies the In predicate on the "check_out_time" field.
func CheckOutTimeIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldCheckOutTime, vs...))
}

// CheckOutTimeNotIn applies the NotIn predicate on the "check_out_time" field.
func CheckOutTimeNotIn(vs ...time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldCheckOutTime, vs...))
}

// CheckOutTimeGT applies the GT predicate on the "check_out_time" field.
func CheckOutTimeGT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGT(FieldCheckOutTime, v))
}

// CheckOutTimeGTE applies the GTE predicate on the "check_out_time" field.
func CheckOutTimeGTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldGTE(FieldCheckOutTime, v))
}

// CheckOutTimeLT applies the LT predicate on the "check_out_time" field.
func CheckOutTimeLT(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLT(FieldCheckOutTime, v))
}

// CheckOutTimeLTE applies the LTE predicate on the "check_out_time" field.
func CheckOutTimeLTE(v time.Time) predicate.Attendance {
	return predicate.Attendance(sql.FieldLTE(FieldCheckOutTime, v))
}

// CheckOutTimeIsNil applies the IsNil predicate on the "check_out_time" field.
func CheckOutTimeIsNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldIsNull(FieldCheckOutTime))
}

// CheckOutTimeNotNil applies the NotNil predicate on the "check_out_time" field.
func CheckOutTimeNotNil() predicate.Attendance {
	return predicate.Attendance(sql.FieldNotNull(FieldCheckOutTime))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Attendance {
	return predicate.Attendance(sql.FieldNotIn(FieldStatus, vs...))
}

// HasEmployee applies the HasEdge predicate on the "employee" edge.
func HasEmployee() predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EmployeeTable, EmployeeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmployeeWith applies the HasEdge predicate on the "employee" edge with a given conditions (other predicates).
func HasEmployeeWith(preds ...predicate.Employee) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		step := newEmployeeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attendance) predicate.Attendance {
	return predicate.Attendance(func(s *sql.Selector) {
		p(s.Not())
	})
}
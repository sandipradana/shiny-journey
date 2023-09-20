package repository

import (
	"context"

	"shiny-journey/ent"
	"shiny-journey/ent/attendance"
)

type AttendanceRepository interface {
	CreateAttendance(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error)
	GetAttendanceByID(ctx context.Context, id int) (*ent.Attendance, error)
	UpdateAttendance(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error)
	DeleteAttendance(ctx context.Context, id int) error
}

type attendanceRepository struct {
	client *ent.Client
}

func NewAttendanceRepository(client *ent.Client) AttendanceRepository {
	return &attendanceRepository{
		client: client,
	}
}

func (r *attendanceRepository) CreateAttendance(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error) {
	createdAttendance, err := r.client.Attendance.
		Create().
		SetAttendanceDate(attendance.AttendanceDate).
		SetNillableCheckInTime(attendance.CheckInTime).
		SetNillableCheckOutTime(attendance.CheckOutTime).
		SetStatus(attendance.Status).
		SetEmployee(attendance.Edges.Employee).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return createdAttendance, nil
}

func (r *attendanceRepository) GetAttendanceByID(ctx context.Context, id int) (*ent.Attendance, error) {
	attendance, err := r.client.Attendance.
		Query().
		Where(attendance.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

func (r *attendanceRepository) UpdateAttendance(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error) {
	updatedAttendance, err := r.client.Attendance.
		UpdateOneID(attendance.ID).
		SetStatus(attendance.Status).
		SetAttendanceDate(attendance.AttendanceDate).
		SetNillableCheckInTime(attendance.CheckInTime).
		SetNillableCheckOutTime(attendance.CheckOutTime).
		SetEmployee(attendance.Edges.Employee).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedAttendance, nil
}

func (r *attendanceRepository) DeleteAttendance(ctx context.Context, id int) error {
	return r.client.Attendance.
		DeleteOneID(id).
		Exec(ctx)
}

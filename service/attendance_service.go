package service

import (
	"context"
	"shiny-journey/ent"
	"shiny-journey/repository"
)

type AttendanceService struct {
	repo repository.AttendanceRepository
}

func NewAttendanceService(repo repository.AttendanceRepository) *AttendanceService {
	return &AttendanceService{
		repo: repo,
	}
}

func (s *AttendanceService) CreateAttendance(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error) {
	return s.repo.CreateAttendance(ctx, attendance)
}

func (s *AttendanceService) GetAttendanceByID(ctx context.Context, id int) (*ent.Attendance, error) {
	return s.repo.GetAttendanceByID(ctx, id)
}

func (s *AttendanceService) UpdateAttendance(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error) {
	return s.repo.UpdateAttendance(ctx, attendance)
}

func (s *AttendanceService) DeleteAttendance(ctx context.Context, id int) error {
	return s.repo.DeleteAttendance(ctx, id)
}

package service

import (
	"context"
	"shiny-journey/ent"
	"shiny-journey/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		repo: repo,
	}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee *ent.Employee) (*ent.Employee, error) {
	return s.repo.CreateEmployee(ctx, employee)
}

func (s *EmployeeService) GetEmployees(ctx context.Context) ([]*ent.Employee, error) {
	return s.repo.GetEmployees(ctx)
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id int) (*ent.Employee, error) {
	return s.repo.GetEmployeeByID(ctx, id)
}

func (s *EmployeeService) GetEmployeeByEmail(ctx context.Context, email string) (*ent.Employee, error) {
	return s.repo.GetEmployeeByEmail(ctx, email)
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, employee *ent.Employee) (*ent.Employee, error) {
	return s.repo.UpdateEmployee(ctx, employee)
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.repo.DeleteEmployee(ctx, id)
}

func (s *EmployeeService) GetEmployeesWithoutClockIn(ctx context.Context) ([]*ent.Employee, error) {
	return s.repo.GetEmployeesWithoutAttendance(ctx, "clock_in")
}

func (s *EmployeeService) GetEmployeesWithoutClockOut(ctx context.Context) ([]*ent.Employee, error) {
	return s.repo.GetEmployeesWithoutAttendance(ctx, "clock_out")
}

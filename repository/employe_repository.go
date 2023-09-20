package repository

import (
	"context"
	"database/sql"
	"fmt"
	"shiny-journey/ent"
	"shiny-journey/ent/employee"
	"time"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *ent.Employee) (*ent.Employee, error)
	GetEmployeeByID(ctx context.Context, id int) (*ent.Employee, error)
	GetEmployees(ctx context.Context) ([]*ent.Employee, error)
	GetEmployeeByEmail(ctx context.Context, email string) (*ent.Employee, error)
	UpdateEmployee(ctx context.Context, employee *ent.Employee) (*ent.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
	GetEmployeesWithoutAttendance(ctx context.Context, t string) ([]*ent.Employee, error)
}
type employeeRepository struct {
	client *ent.Client
	db     *sql.DB
}

func NewEmployeeRepository(client *ent.Client, db *sql.DB) EmployeeRepository {
	return &employeeRepository{
		client: client,
		db:     db,
	}
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee *ent.Employee) (*ent.Employee, error) {
	createdEmployee, err := r.client.Employee.
		Create().
		SetName(employee.Name).
		SetEmail(employee.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return createdEmployee, nil
}

func (r *employeeRepository) GetEmployeeByID(ctx context.Context, id int) (*ent.Employee, error) {
	employee, err := r.client.Employee.
		Query().
		Where(employee.IDEQ(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *employeeRepository) GetEmployees(ctx context.Context) ([]*ent.Employee, error) {
	employees, err := r.client.Employee.
		Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *employeeRepository) GetEmployeeByEmail(ctx context.Context, email string) (*ent.Employee, error) {
	employee, err := r.client.Employee.
		Query().
		Where(employee.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *employeeRepository) UpdateEmployee(ctx context.Context, employee *ent.Employee) (*ent.Employee, error) {
	updatedEmployee, err := r.client.Employee.
		UpdateOneID(employee.ID).
		SetName(employee.Name).
		SetEmail(employee.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedEmployee, nil
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, id int) error {
	return r.client.Employee.
		DeleteOneID(id).
		Exec(ctx)
}

func (r *employeeRepository) GetEmployeesWithoutAttendance(ctx context.Context, t string) ([]*ent.Employee, error) {

	addQuery := ""

	if t == "clock_in" {
		addQuery = "AND a.check_in_time IS NULL"
	}

	if t == "clock_out" {
		addQuery = "AND a.check_out_time IS NULL"
	}

	query := fmt.Sprintf(`SELECT e.* FROM employees e LEFT JOIN attendances a ON e.id = a.employee_attendances WHERE a.attendance_date = "%s" %s`, time.Now().Format("2006-01-02"), addQuery)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*ent.Employee
	for rows.Next() {
		var employee ent.Employee
		if err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
		); err != nil {
			return nil, err
		}

		employees = append(employees, &employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

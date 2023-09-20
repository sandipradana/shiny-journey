package handler

import (
	"context"
	"net/http"
	"shiny-journey/ent"
	"shiny-journey/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	employeeService *service.EmployeeService
}

func NewEmployeeHandler(employeeService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}

func (h *EmployeeHandler) RegisterRoutes(e *echo.Echo, jwtMiddleware echo.MiddlewareFunc) {
	e.POST("/employees", h.CreateEmployee, jwtMiddleware)
	e.GET("/employees/:id", h.GetEmployee, jwtMiddleware)
	e.PUT("/employees/:id", h.UpdateEmployee, jwtMiddleware)
	e.DELETE("/employees/:id", h.DeleteEmployee, jwtMiddleware)
}

func (h *EmployeeHandler) CreateEmployee(c echo.Context) error {
	ctx := context.Background()

	var employeeRequest ent.Employee
	if err := c.Bind(&employeeRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	createdEmployee, err := h.employeeService.CreateEmployee(ctx, &employeeRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Employee creation failed"})
	}

	return c.JSON(http.StatusCreated, createdEmployee)
}

func (h *EmployeeHandler) GetEmployee(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	employee, err := h.employeeService.GetEmployeeByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Employee not found"})
	}

	return c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	existingEmployee, err := h.employeeService.GetEmployeeByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Employee not found"})
	}

	var newEmployee ent.Employee
	if err := c.Bind(&newEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	newEmployee.ID = existingEmployee.ID
	updatedEmployee, err := h.employeeService.UpdateEmployee(ctx, &newEmployee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Employee update failed"})
	}

	return c.JSON(http.StatusOK, updatedEmployee)
}

func (h *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid employee ID"})
	}

	err = h.employeeService.DeleteEmployee(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Employee deletion failed"})
	}

	return c.NoContent(http.StatusNoContent)
}

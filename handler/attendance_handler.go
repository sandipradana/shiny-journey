package handler

import (
	"context"
	"net/http"
	"shiny-journey/ent"
	"shiny-journey/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AttendanceHandler struct {
	attendanceService *service.AttendanceService
}

func NewAttendanceHandler(attendanceService *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

func (h *AttendanceHandler) RegisterRoutes(e *echo.Echo, jwtMiddleware echo.MiddlewareFunc) {
	e.POST("/attendances", h.CreateAttendance, jwtMiddleware)
	e.GET("/attendances/:id", h.GetAttendance, jwtMiddleware)
	e.PUT("/attendances/:id", h.UpdateAttendance, jwtMiddleware)
	e.DELETE("/attendances/:id", h.DeleteAttendance, jwtMiddleware)
}

func (h *AttendanceHandler) CreateAttendance(c echo.Context) error {
	ctx := context.Background()

	var attendanceRequest ent.Attendance
	if err := c.Bind(&attendanceRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	createdAttendance, err := h.attendanceService.CreateAttendance(ctx, &attendanceRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Attendance creation failed"})
	}

	return c.JSON(http.StatusCreated, createdAttendance)
}

func (h *AttendanceHandler) GetAttendance(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid attendance record ID"})
	}

	attendance, err := h.attendanceService.GetAttendanceByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Attendance record not found"})
	}

	return c.JSON(http.StatusOK, attendance)
}

func (h *AttendanceHandler) UpdateAttendance(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid attendance record ID"})
	}

	existingAttendance, err := h.attendanceService.GetAttendanceByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Attendance record not found"})
	}

	var newAttendance ent.Attendance
	if err := c.Bind(&newAttendance); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	newAttendance.ID = existingAttendance.ID
	updatedAttendance, err := h.attendanceService.UpdateAttendance(ctx, &newAttendance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Attendance record update failed"})
	}

	return c.JSON(http.StatusOK, updatedAttendance)
}

func (h *AttendanceHandler) DeleteAttendance(c echo.Context) error {
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid attendance record ID"})
	}

	err = h.attendanceService.DeleteAttendance(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Attendance record deletion failed"})
	}

	return c.NoContent(http.StatusNoContent)
}

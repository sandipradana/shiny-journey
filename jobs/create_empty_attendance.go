package jobs

import (
	"context"
	"fmt"
	"shiny-journey/ent"
	"shiny-journey/service"
	"time"
)

func createEmptyAttendance(employeeService *service.EmployeeService, attendanceService *service.AttendanceService) func() {
	return func() {
		employees, err := employeeService.GetEmployees(context.Background())
		if err != nil {
			return
		}

		for i := range employees {
			attendance, err := attendanceService.CreateAttendance(context.Background(), &ent.Attendance{
				AttendanceDate: time.Now(),
				Status:         "absent",
				Edges: ent.AttendanceEdges{
					Employee: employees[i],
				},
			})
			if err != nil {
				return
			}
			fmt.Println(attendance)
		}
	}
}

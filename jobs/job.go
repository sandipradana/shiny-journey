package jobs

import (
	"os"
	"os/signal"
	"shiny-journey/service"
	"syscall"
	"time"

	"github.com/robfig/cron/v3"
)

func New(employeeService *service.EmployeeService, attendanceService *service.AttendanceService) {
	scheduler := cron.New(cron.WithLocation(time.Local))
	defer scheduler.Stop()

	scheduler.AddFunc("0 0 * * *", createEmptyAttendance(employeeService, attendanceService))
	scheduler.AddFunc("45 8 * * *", clockReminder(employeeService, "clock_in"))
	scheduler.AddFunc("45 16 * * *", clockReminder(employeeService, "clock_out"))

	go scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

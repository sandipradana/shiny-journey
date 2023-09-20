package jobs

import (
	"context"
	"fmt"
	"os"
	"shiny-journey/ent"
	"shiny-journey/service"

	"github.com/go-gomail/gomail"
)

func clockReminder(employeeService *service.EmployeeService, t string) func() {
	return func() {
		if t == "clock_in" {
			employees, err := employeeService.GetEmployeesWithoutClockIn(context.Background())
			if err != nil {
				fmt.Println(err)
				return
			}
			for i := range employees {
				sendEmailReminder(employees[i], "Clock In")
			}
		} else if t == "clock_out" {
			employees, err := employeeService.GetEmployeesWithoutClockOut(context.Background())
			if err != nil {
				fmt.Println(err)
				return
			}
			for i := range employees {
				sendEmailReminder(employees[i], "Clock Out")
			}
		}
	}
}

func sendEmailReminder(employee *ent.Employee, event string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "mail@sandipradana.com")
	mail.SetHeader("To", employee.Email)
	mail.SetHeader("Subject", fmt.Sprintf("Reminder: %s", event))
	mail.SetBody("text/plain", fmt.Sprintf("Don't forget to %s soon!", event))

	d := gomail.NewDialer("smtp.zoho.com", 465, "mail@sandipradana.com", os.Getenv("MAIL_PASSWORD"))

	if err := d.DialAndSend(mail); err != nil {
		fmt.Printf("Failed to send email reminder: %v\n", err)
	}
}

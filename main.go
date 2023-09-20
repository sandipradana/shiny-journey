package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"shiny-journey/ent"
	"shiny-journey/handler"
	"shiny-journey/jobs"
	"shiny-journey/repository"
	"shiny-journey/service"
	"syscall"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

const jwtSecret = "JwtSecret321!"

func main() {

	drv, err := sql.Open("mysql", "root@tcp(localhost:3306)/attendances")
	if err != nil {
		log.Fatalf("%v", err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	adminRepository := repository.NewAdminRepository(client)
	employeeRepository := repository.NewEmployeeRepository(client, db)
	attendanceRepository := repository.NewAttendanceRepository(client)

	authService := service.NewAuthService(adminRepository, jwtSecret)
	employeeService := service.NewEmployeeService(employeeRepository)
	attendanceService := service.NewAttendanceService(attendanceRepository)

	authHandler := handler.NewAuthHandler(authService)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	attendaceHandler := handler.NewAttendanceHandler(attendanceService)

	e := echo.New()

	authHandler.RegisterRoutes(e)
	employeeHandler.RegisterRoutes(e)
	attendaceHandler.RegisterRoutes(e)

	go e.Start(":8080")
	go jobs.New(employeeService, attendanceService)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

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
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

const jwtSecret = "JwtSecret321!"

func main() {

	drv, err := sql.Open("sqlite3", "file:ent.sqlite3?mode=memory&cache=shared&_fk=1")
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

	jwtMiddleware := echojwt.JWT([]byte(jwtSecret))

	authHandler.RegisterRoutes(e, jwtMiddleware)
	employeeHandler.RegisterRoutes(e, jwtMiddleware)
	attendaceHandler.RegisterRoutes(e, jwtMiddleware)

	go e.Start(":8080")
	go jobs.New(employeeService, attendanceService)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

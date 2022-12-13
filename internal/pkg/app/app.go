package app

import (
	"ServiceR2/internal/app/endpoint"
	"ServiceR2/internal/app/middleware"
	"ServiceR2/internal/app/service"
	"log"

	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleware.ChekUser)

	a.echo.GET("/", nil)
	a.echo.GET("/fibonacci/:number", a.e.Fibonacci).Name = "number"
	a.echo.GET("/time-moscow", a.e.TimeInMoscow)

	return a, nil

}

func (a *App) Run() error {
	log.Println("start service...")

	if err := a.echo.Start(":8080"); err != nil {
		log.Fatal("service return error...", err)
	}

	return nil

}

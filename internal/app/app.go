package app

import (
	"Breeding/internal/app/endpoint"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	ep   *endpoint.Endpoint
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.ep = endpoint.New()
	a.echo = echo.New()
	a.echo.POST("/get_pdf", a.ep.Post)
	a.echo.File("/", a.ep.IndexPage)

	return a, nil
}

func (s *App) Run() error {
	log.Println("Service running")
	err := s.echo.Start(":80")
	if err != nil {
		return err
	}
	return nil
}

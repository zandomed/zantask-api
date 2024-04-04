package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zandomed/zantask-api/config"
	"github.com/zandomed/zantask-api/database"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.loadRoutes()

	port := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(port))
}

func (s *echoServer) Stop() {
	s.app.Close()
}

func (s *echoServer) loadRoutes() {
	// Load the handler here

	s.app.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// v1 := s.app.Group("/v1")

}

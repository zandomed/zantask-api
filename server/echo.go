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
	App  *echo.Echo
	DB   database.Database
	Conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		App:  echoApp,
		DB:   db,
		Conf: conf,
	}
}

func (s *echoServer) Start() {
	s.App.Use(middleware.Recover())
	s.App.Use(middleware.Logger())

	s.loadRoutes()

	port := fmt.Sprintf(":%d", s.Conf.Server.Port)
	s.App.Logger.Fatal(s.App.Start(port))
}

func (s *echoServer) Stop() {
	s.App.Close()
}

func (s *echoServer) loadRoutes() {
	// Load the handler here

	s.App.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// v1 := s.app.Group("/v1")

}

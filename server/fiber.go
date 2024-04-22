package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/zandomed/zantask-api/config"
	"github.com/zandomed/zantask-api/database"
)

type fiberServer struct {
	App  *fiber.App
	DB   database.Database
	Conf *config.Config
}

func NewFiberServer(conf *config.Config, db database.Database) Server {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "ZanTask",
		AppName:       "ZanTask",
	})
	log.SetLevel(log.LevelDebug)

	app.Use(helmet.New())
	app.Use(healthcheck.NewHealthChecker())

	return &fiberServer{
		App:  app,
		DB:   db,
		Conf: conf,
	}
}

func (s *fiberServer) Start() {

	s.loadRoutes()

	port := fmt.Sprintf(":%d", s.Conf.Server.Port)
	log.Fatal(s.App.Listen(port))
}

func (s *fiberServer) Stop() {
	if err := s.App.Shutdown(); err != nil {
		panic(fmt.Sprintf("Error while shutting down the server: %v", err))
	}
}

func (s *fiberServer) loadRoutes() {
	// Load the handler here

	s.App.Get("/", func(c fiber.Ctx) error {
		return c.JSON(map[string]string{"status": "OK"})
	})
}

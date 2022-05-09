package api

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

type Log struct {
	Data string `json:"data"`
}

func CreateRoutes(server *fiber.App) {
	/// Prometheus middleware
	prometheus := fiberprometheus.New("hades")
	prometheus.RegisterAt(server, "/metrics")
	server.Use(prometheus.Middleware)

	// Display all logs
	server.Get("/logs", Index)
	// Get log by id
	server.Get("/logs/:id", ByID)
	// Publish log
	server.Post("/publish", Create)
}

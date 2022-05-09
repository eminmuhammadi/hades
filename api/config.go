package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Docs: https://docs.gofiber.io
var APP_CONFIGURATION = fiber.Config{
	EnablePrintRoutes:     true,
	ReduceMemoryUsage:     true,
	BodyLimit:             4 * 1024 * 1024,
	ReadTimeout:           time.Second * 15,
	WriteTimeout:          time.Second * 15,
	IdleTimeout:           time.Second * 60,
	ReadBufferSize:        4096,
	WriteBufferSize:       4096,
	DisableStartupMessage: false,
	AppName:               "Hades v1.0",
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(500).SendString(err.Error())
}

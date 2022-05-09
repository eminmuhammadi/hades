package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// CreateServer creates a new server instance
func CreateServer() *fiber.App {
	return fiber.New(APP_CONFIGURATION)
}

// Creates a listener
func StartServer(app *fiber.App) error {
	log.Printf("%s has been started", APP_CONFIGURATION.AppName)
	log.Printf(
		"%s is listening on port %s:%s",
		APP_CONFIGURATION.AppName,
		os.Getenv("HOSTNAME"),
		os.Getenv("PORT"),
	)

	return app.Listen(fmt.Sprintf(
		"%s:%s",
		os.Getenv("HOSTNAME"),
		os.Getenv("PORT"),
	))
}

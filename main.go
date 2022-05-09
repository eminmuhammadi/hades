package main

import (
	"fmt"
	"os"
	"os/signal"

	api "github.com/eminmuhammadi/hades/api"
	broker "github.com/eminmuhammadi/hades/broker"
	data "github.com/eminmuhammadi/hades/data"
	"github.com/eminmuhammadi/hades/db"
)

// Migrate the database on startup
func Setup() {
	sql := db.Open()

	db.AutoMigrate(sql)

	defer db.Close(sql)
}

func main() {
	Setup()

	// Graceful shutdown configuration
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	serverShutdown := make(chan struct{})

	// Create the Server
	app := api.CreateServer()

	// Graceful shutdown prompt
	go func() {
		<-quit
		fmt.Println("Gracefully shutting down...")
		app.Shutdown()
		serverShutdown <- struct{}{}
	}()

	// Routes
	api.CreateRoutes(app)

	// AMPQ Consumer
	consumer := broker.ConsumerStart()
	data.ConsumeData(consumer)

	// Start server
	if err := api.StartServer(app); err != nil {
		panic(err)
	}

	<-serverShutdown

	// Cleanup tasks
	defer consumer.Close()
}

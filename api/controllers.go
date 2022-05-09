package api

import (
	broker "github.com/eminmuhammadi/hades/broker"
	data "github.com/eminmuhammadi/hades/data"
	db "github.com/eminmuhammadi/hades/db"
	model "github.com/eminmuhammadi/hades/model"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

// Display all logs
func Index(ctx *fiber.Ctx) error {
	sql := db.Open()
	_model := sql.Model(&model.Log{})

	_paginate := paginate.New()
	_paginate.Config.SmartSearch = true

	// ref github.com/morkid/paginate
	data := _paginate.With(_model).
		Request(ctx.Request()).
		Response(&[]model.Log{})

	ctx.JSON(&fiber.Map{
		"success": true,
		"error":   nil,
		"data":    data,
	})

	defer db.Close(sql)
	return nil
}

// Get log by id
func ByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	log, err := data.GetLogByID(id)

	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"error":   "Log not found",
			"data":    nil,
		})
	}

	ctx.Status(200).JSON(&fiber.Map{
		"success": true,
		"error":   nil,
		"data":    log,
	})

	return nil
}

// Publish log
func Create(ctx *fiber.Ctx) error {
	log := new(Log)

	if err := ctx.BodyParser(log); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"sucess": false,
			"error":  err.Error(),
			"data":   nil,
		})

		return err
	}

	if log.Data == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"sucess": false,
			"error":  "Data is empty",
			"data":   nil,
		})

		return nil
	}

	publisher := broker.PublisherStart()
	broker.Publish([]byte(log.Data), []string{"logcollector"}, publisher)

	ctx.Status(201).JSON(&fiber.Map{
		"success": true,
		"error":   nil,
		"data":    nil,
	})

	defer publisher.Close()
	return nil
}

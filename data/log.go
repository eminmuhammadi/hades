package data

import (
	"os"

	broker "github.com/eminmuhammadi/hades/broker"
	db "github.com/eminmuhammadi/hades/db"
	model "github.com/eminmuhammadi/hades/model"
	"github.com/wagslane/go-rabbitmq"
)

// Create log
func CreateLog(log model.Log) error {
	sql := db.Open()

	if err := sql.Create(&log).Error; err != nil {
		return err
	}

	defer db.Close(sql)
	return nil
}

// Get log by id
func GetLogByID(id string) (model.Log, error) {
	sql := db.Open()

	var log model.Log
	if err := sql.Where("id = ?", id).First(&log).Error; err != nil {
		return model.Log{}, err
	}

	defer db.Close(sql)
	return log, nil
}

// Consume data from rabbitmq
func ConsumeData(consumer rabbitmq.Consumer) rabbitmq.Consumer {
	return broker.Consume(
		func(data rabbitmq.Delivery) rabbitmq.Action {
			sql := db.Open()

			log := model.Log{
				Data: string(data.Body),
			}

			if err := CreateLog(log); err != nil {
				return rabbitmq.NackDiscard
			}

			defer db.Close(sql)
			return rabbitmq.Ack
		},
		os.Getenv("QUEUE"),
		[]string{"logcollector"},
		consumer,
	)
}

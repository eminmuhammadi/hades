package broker

import (
	"github.com/wagslane/go-rabbitmq"
)

// PublisherStart starts the publisher
func PublisherStart() *rabbitmq.Publisher {
	publisher, err := rabbitmq.NewPublisher(AMPQ_URL, rabbitmq.Config{},
		rabbitmq.WithPublisherOptionsLogging,
	)

	// if there is an error, panic
	if err != nil {
		panic(err)
	}

	return publisher
}

// Publisher is a wrapper for the RabbitMQ publisher
func Publish(
	message []byte,
	route []string,
	publisher *rabbitmq.Publisher,
) *rabbitmq.Publisher {
	if err := publisher.Publish(
		message,
		route,
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsContentEncoding("utf-8"),
		rabbitmq.WithPublishOptionsExchange("amq.direct"),
	); err != nil {
		panic(err)
	}

	return publisher
}

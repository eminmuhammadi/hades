package broker

import (
	"github.com/wagslane/go-rabbitmq"
)

// ConsumerStart starts the consumer
func ConsumerStart() rabbitmq.Consumer {
	consumer, err := rabbitmq.NewConsumer(AMPQ_URL, rabbitmq.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)

	// if there is an error, panic
	if err != nil {
		panic(err)
	}

	return consumer
}

// Consumer is a wrapper for the RabbitMQ consumer
func Consume(
	handler rabbitmq.Handler,
	queue string,
	route []string,
	consumer rabbitmq.Consumer,
) rabbitmq.Consumer {
	// Start consuming messages
	if err := consumer.StartConsuming(
		handler,
		queue,
		route,
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsQuorum,
		rabbitmq.WithConsumeOptionsBindingExchangeDurable,
		rabbitmq.WithConsumeOptionsBindingExchangeName("amq.direct"),
		rabbitmq.WithConsumeOptionsBindingExchangeKind("direct"),
		rabbitmq.WithConsumeOptionsConsumerName("hades-consumer"),
	); err != nil {
		panic(err)
	}

	return consumer
}

package broker

import (
	"fmt"
	"os"
)

var AMPQ_URL = fmt.Sprintf(
	"amqp://%s:%s@%s:%s/%s",
	os.Getenv("AMPQ_USERNAME"),
	os.Getenv("AMPQ_PASSWORD"),
	os.Getenv("AMPQ_HOST"),
	os.Getenv("AMPQ_PORT"),
	os.Getenv("AMPQ_VHOST"),
)

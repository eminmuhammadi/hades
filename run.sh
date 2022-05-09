# /bin/bash

# Config
export TIMEZONE=UTC
export PORT=8080
export HOSTNAME=localhost
export QUEUE=hades_queue

# RabbitMQ
# Free at https://www.cloudamqp.com
export AMPQ_USERNAME=
export AMPQ_PASSWORD=
export AMPQ_HOST=
export AMPQ_VHOST=

# PostgreSQL
# Free at https://www.elephantsql.com
export DB_HOST=
export DB_USERNAME=
export DB_PASSWORD=
export DB_NAME=
export DB_PORT=5432
export DB_SSLMODE=require

go run main.go
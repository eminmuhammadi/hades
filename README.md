# Hades

Hades is a simple, fast and continuous log collector.

`SERVER` => `HADES_API` => `PUBLISHER` => `RABBITMQ` => `CONSUMER` => `POSTGRES`

## Development
```bash
bash run.sh
```

## Production
docker-compose up --build -d
```

## Controllers
```
method  | path      | name | handlers
------  | ----      | ---- | --------
POST    | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
GET     | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
TRACE   | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
OPTIONS | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
CONNECT | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
HEAD    | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
DELETE  | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
PUT     | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
PATCH   | /         |      | github.com/ansrivas/fiberprometheus/v2.(*FiberPrometheus).Middleware-fm
HEAD    | /logs     |      | github.com/eminmuhammadi/hades/api.Index
GET     | /logs     |      | github.com/eminmuhammadi/hades/api.Index
HEAD    | /logs/:id |      | github.com/eminmuhammadi/hades/api.ByID
GET     | /logs/:id |      | github.com/eminmuhammadi/hades/api.ByID
GET     | /metrics  |      | github.com/gofiber/adaptor/v2.HTTPHandler.func1
HEAD    | /metrics  |      | github.com/gofiber/adaptor/v2.HTTPHandler.func1
POST    | /publish  |      | github.com/eminmuhammadi/hades/api.Create
```

/publish is a POST request that accepts a JSON payload. The payload should be a map with the following keys:
- `data`: the log data
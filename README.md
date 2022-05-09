# Hades

Hades is a simple, fast and continuous log collector.

`SERVER` => `HADES_API` => `PUBLISHER` => `RABBITMQ` => `CONSUMER` => `POSTGRES`

## Development
```bash
bash run.sh
```

## Production
```bash
docker-compose up --build -d
```

## Controllers
```bash
method  | path      | name | handlers
------  | ----      | ---- | --------
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

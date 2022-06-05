
## Stack:
- Golang
- GraphQ: [gqlgen](github.com/99designs/gqlgen)
- PostgreSQL
- Migrations: [dbmate](https://github.com/amacneil/dbmate)
- mocks for tests - [moq](github.com/matryer/moq)
- [reflex](https://github.com/cespare/reflex) for server reload at editing or crushes
- RabbitMQ

## Setup:
- Install and setup [golang](https://go.dev/)
- `git clone https://github.com/Spargwy/transactions`
- `cd transactions`
- `go mod tidy`

## Run:
- `docker-compose up -d postgres`
- `dbmate up`
- `make apply` - create user(its possible from api)
- `cd client && ./scripts dev`
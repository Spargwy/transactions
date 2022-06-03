
## Stack:
- Golang
- GraphQ: [gqlgen](github.com/99designs/gqlgen)
- PostgreSQL
- Migrations: [dbmate](https://github.com/amacneil/dbmate)
- mocks for tests - [moq](github.com/matryer/moq)
- [reflex](https://github.com/cespare/reflex) for server reload at editing or crushes
- RabbitMQ

## Run:
- `docker-compose up -d postgres`
- `dbmate up`
-  `make apply`
- `cd client && ./scripts dev`

package env

import (
	"transactions/utils"

	"github.com/go-pg/pg/v10"
)

//All tools that we want to use
type Env struct {
	db *pg.DB
}

const defaultConnectionURL = "postgres://staging:staging@localhost/staging?sslmode=disable"

func New() *Env {
	appEnv := utils.GetEnv("APP_MODE", "development")

	DB := dbConnect(utils.GetEnv("DB_CONNECTION_STRING", defaultConnectionURL), appEnv)

	env := Env{
		db: DB,
	}

	return &env
}

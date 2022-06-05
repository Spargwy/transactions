package env

import (
	"log"
	"transactions/utils"

	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"
)

//Необходимые тулы
type Env struct {
	db    *pg.DB
	rmq   *amqp.Connection
	rmqCh *amqp.Channel
}

const defaultConnectionURL = "postgres://staging:staging@localhost/staging?sslmode=disable"

func New() *Env {
	DB := dbConnect(utils.GetEnv("DB_CONNECTION_STRING", defaultConnectionURL))
	rmq, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	rmqCh, err := rmq.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}

	log.Print("Connected to RabbitMQ")

	env := Env{
		db:    DB,
		rmq:   rmq,
		rmqCh: rmqCh,
	}

	users, err := env.GetAllUsers()
	if err != nil {
		log.Fatalf("failed to GetAllUsers: %v", err)
	}

	for _, user := range users {
		//При старте(или рестарте) приложения
		//По новой создаются очереди и слушатели очередей
		//для каждого пользователя
		q, err := env.CreateUserQueue(user.ID)
		if err != nil {
			log.Fatalf("failed to CreateUserQueue: %v", err)
		}

		go env.CreateConsumerForQueue(q.Name)
	}

	return &env
}

package env

import (
	"fmt"
	"log"
	"strconv"
	"transactions/model"
	"transactions/service/processTransaction"

	"github.com/streadway/amqp"
)

func (e *Env) CreateUserQueue(userID int) (amqp.Queue, error) {
	//Очередь для каждого пользователя
	q, err := e.rmqCh.QueueDeclare(strconv.Itoa(userID),
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		return amqp.Queue{}, fmt.Errorf("failed to QueueDeclare: %v", err)
	}

	return q, err
}

func (e *Env) PublishRMQMessage(qName string, body []byte) error {
	err := e.rmqCh.Publish(
		"",
		qName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return err
}

func (e *Env) CreateConsumerForQueue(qname string) {
	msgs, err := e.rmqCh.Consume(
		qname,
		qname,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("failed to create Consume: %v", err)
		return
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			transaction, err := processTransaction.Resolve(e, msg)
			if err != nil {
				//Чтобы не было копипаста, обрабатываем фейлы здесь
				transaction.Status = model.FailedOperationStatus
				transaction.FinishedAt = e.Now()
				err = e.UpdateTransaction(&transaction)
				if err != nil {
					log.Printf("failed to UpdateTransaction: %v", err)
				}
			}
		}
	}()

	<-forever
}

package broker

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Broker struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func NewBroker(user, passwod, port string) *Broker {
	broker := &Broker{}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@rabbit:%s",
		user,
		passwod,
		port))
	failOnError(err, "failed to connect")
	ch, err := conn.Channel()
	failOnError(err, "failed to open channel")
	q, err := ch.QueueDeclare(
		viper.GetString("queue_name"),
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare")
	broker.Channel = ch
	broker.Queue = q
	return broker
}

func (s *Broker) Publish(message interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	parsedMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = s.Channel.PublishWithContext(ctx,
		"",
		s.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        parsedMessage,
		})
	if err != nil {
		return err
	}
	log.Println(" [x] Sent ", message)
	return nil
}

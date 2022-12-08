package main

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"log"
)

func main() {
	client := NewClient()
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@rabbit:%s",
		viper.GetString("rabbit_user"),
		viper.GetString("rabbit_password"),
		viper.GetString("rabbit_pot")))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		viper.GetString("queue_name"), // name
		false,                         // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			var result GithubRepository
			err := json.Unmarshal(d.Body, &result)
			if err != nil {
				client.SendRequestToGithubApi(fmt.Sprint(d.Body))
			} else {
				client.SendRequestToMainService(result)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func initConfig() error {
	viper.AddConfigPath("web/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

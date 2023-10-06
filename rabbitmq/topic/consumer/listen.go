package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

const _address = "amqp://1:1@localhost:5672/"

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	conn, err := amqp.Dial(_address)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		err = conn.Close()
		failOnError(err, "Failed to close connect to RabbitMQ")
	}()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() {
		err = ch.Close()
		failOnError(err, "Failed to close a channel")
	}()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "usless-topic", s)
		err = ch.QueueBind(
			q.Name,         // queue name
			s,              // routing key
			"usless-topic", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}

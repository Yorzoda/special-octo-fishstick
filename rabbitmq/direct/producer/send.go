package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strings"
	"time"
)

const _address = "amqp://1:1@localhost:5672/"

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
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

	err = ch.ExchangeDeclare(
		"simple.topic",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a Exchange")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(ctx,
		"simple.topic", // exchange
		"",             // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}

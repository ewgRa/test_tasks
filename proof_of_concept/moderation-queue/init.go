package main

import (
	"log"
	"github.com/streadway/amqp"
)


func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"moderation.category1", // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	_, err = ch.QueueDeclare(
		"moderation.category2", // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		nil,                    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.ExchangeDeclare("moderation", "topic", true, // durable
		false,                                            // delete when unused
		false,                                            // exclusive
		false,                                            // no-wait
		nil)
	failOnError(err, "Failed to open a channel")

	err = ch.QueueBind("moderation.category1", "#.category1.#", "moderation", false, nil)
	failOnError(err, "Failed to open a channel")

	err = ch.QueueBind("moderation.category2", "#.category2.#", "moderation", false, nil)
	failOnError(err, "Failed to open a channel")
}

package main

import (
	"flag"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var objId  = flag.String("id", "", "Object id, eg. 1, 2")
	flag.Parse()

	if *objId == "" {
		flag.Usage()
		os.Exit(1)
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Publish(
		"moderation",     // exchange
		"moderation.foo.category1.bar.category2.foobar", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(*objId),
		})
	log.Printf(" [x] Sent %s", *objId)
	failOnError(err, "Failed to publish a message")
}

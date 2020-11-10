package main

import (
	"log"
	"flag"
	"os"
	"time"

	"github.com/streadway/amqp"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var category  = flag.String("category", "", "Category name, eg. category1, category2")
	flag.Parse()

	if *category == "" {
		flag.Usage()
		os.Exit(1)
	}

	redisClient := redis.NewClient(&redis.Options{
		Network:	"tcp",
		Addr:		"localhost:6379",
	})
	defer redisClient.Close()

	// Create a new lock client.
	redisLocker := redislock.New(redisClient)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"moderation."+*category, // queue
		"moderation."+*category,     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			lock, err := redisLocker.Obtain(string(d.Body), 3000*time.Millisecond, nil)
			if err == redislock.ErrNotObtained {
				log.Printf("%s is locked, skip processing\n", d.Body)
				return
			}

			failOnError(err, "Failed to lock record in redis")

			defer lock.Release()

			log.Printf("Process message %s\n", d.Body)

			v := redisClient.Get("object_"+string(d.Body))

			if v.Val() == "1" {
				// prevent case when message delivered to several queues, but consuming of some queues postponed
				log.Printf("%s already processed, skip processing\n", d.Body)
			} else {
				redisClient.Set("object_"+string(d.Body), "1", 0)
				time.Sleep(2000*time.Millisecond)
				log.Printf("Message %s processed\n", d.Body)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

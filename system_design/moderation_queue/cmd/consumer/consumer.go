package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	topic, categories := parseArgs()

	if !consume(topic, categories) {
		os.Exit(1)
	}
}

func consume(topic string, categories []string) bool {
	redisClient := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    os.Getenv("REDIS_ADDR"),
	})

	defer closeIt(redisClient)
	redisLocker := redislock.New(redisClient)

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Error().Caller().Err(err).Msg("Failed to establish connection to RabbitMQ")

		return false
	}

	defer closeIt(conn)

	ch, err := conn.Channel()
	if err != nil {
		log.Error().Caller().Err(err).Msg("Failed to open channel")

		return false
	}

	defer closeIt(ch)

	deliveries, err := createConsumers(ch, topic, categories)
	if err != nil {
		log.Error().Caller().Err(err).Msg("Failed to create consumers")

		return false
	}

	forever := make(chan bool)

	for _, msgs := range deliveries {
		go func(msgs <-chan amqp.Delivery) {
			for message := range msgs {
				go processMessage(redisClient, redisLocker, message)
			}
		}(msgs)
	}

	log.Info().Caller().Msg("Waiting for messages")
	<-forever

	return true
}

func createConsumers(ch *amqp.Channel, topic string, categories []string) ([]<-chan amqp.Delivery, error) {
	consumers := make([]<-chan amqp.Delivery, len(categories))

	for _, category := range categories {
		msgs, err := ch.Consume(
			topic+"."+category,
			uuid.New().String(),
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to register consumer: %w", err)
		}

		consumers = append(consumers, msgs)
	}

	return consumers, nil
}

func parseArgs() (string, []string) {
	topic := flag.String("topic", "", "Topic name, eg. test_topic")
	categories := flag.String("categories", "", "Category names, comma-separated, eg. test_category1,test_category2")
	flag.Parse()

	if *topic == "" {
		flag.Usage()

		os.Exit(1)
	}

	return *topic, strings.Split(*categories, ",")
}

func closeIt(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Error().Caller().Err(err).Msg("Failed to close")
	}
}

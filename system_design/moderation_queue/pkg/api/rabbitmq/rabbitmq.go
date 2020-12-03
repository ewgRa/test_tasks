// Package rabbitmq is a high-level wrapper under rabbitmq client.
package rabbitmq

import (
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

// NewRabbitMq creates new RabbitMq instance.
func NewRabbitMq(url string) *RabbitMq {
	return &RabbitMq{url: url}
}

// RabbitMq wrapper..
type RabbitMq struct {
	url string
}

// CreateTopic creates topic.
func (rmq *RabbitMq) CreateTopic(name string) error {
	conn, ch, err := rmq.connection()
	if err != nil {
		return fmt.Errorf("can't establish connection to RabbitMQ: %w", err)
	}

	defer rmq.close(conn, ch)

	return ch.ExchangeDeclare(name, "topic",
		true,
		false,
		false,
		false,
		nil)
}

// CreateCategory creates category.
func (rmq *RabbitMq) CreateCategory(topic string, name string) error {
	conn, ch, err := rmq.connection()
	if err != nil {
		return fmt.Errorf("can't establish connection to RabbitMQ: %w", err)
	}

	defer rmq.close(conn, ch)

	categoryQueue := topic + "." + name

	_, err = ch.QueueDeclare(
		categoryQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("can't declare queue: %w", err)
	}

	return ch.QueueBind(categoryQueue, "#."+name+".#", topic, false, nil)
}

// CreateMessage creates message in specific topic.
func (rmq *RabbitMq) CreateMessage(topic string, message string, categories []string) error {
	conn, ch, err := rmq.connection()
	if err != nil {
		return fmt.Errorf("can't establish connection to RabbitMQ: %w", err)
	}

	defer rmq.close(conn, ch)

	categoriesKey := ""

	if len(categories) > 0 {
		categoriesKey = "." + strings.Join(categories, ".")
	}

	err = ch.Publish(
		topic,
		topic+categoriesKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   uuid.New().String(),
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("can't publish message to topic: %w", err)
	}

	return nil
}

// Ping returns RabbitMQ server availability.
func (rmq *RabbitMq) Ping() (bool, error) {
	conn, ch, err := rmq.connection()
	if err != nil {
		return false, fmt.Errorf("can't establish connection to RabbitMQ: %w", err)
	}

	defer rmq.close(conn, ch)

	return true, nil
}

func (rmq *RabbitMq) connection() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rmq.url)
	if err != nil {
		return nil, nil, fmt.Errorf("can't dial RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		if closeErr := conn.Close(); closeErr != nil {
			log.Error().Caller().Err(closeErr).Msg("Can't close RabbitMQ connection")
		}

		return nil, nil, fmt.Errorf("can't create RabbitMQ channel: %w", err)
	}

	return conn, ch, nil
}

func (rmq *RabbitMq) close(conn io.Closer, ch io.Closer) {
	if err := conn.Close(); err != nil {
		log.Error().Caller().Err(err).Msg("Can't close RabbitMQ connection")
	}

	if err := ch.Close(); err != nil {
		log.Error().Caller().Err(err).Msg("Can't close RabbitMQ channel")
	}
}

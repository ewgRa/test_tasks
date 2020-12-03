package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

const (
	// MessageTTL represents moderation time limit. It is a time that moderator have for moderate message.
	MessageTTL = 3000 * time.Millisecond
	// Processed represents a flag value for message that marked as processed.
	Processed = "processed"
)

func processMessage(redisClient *redis.Client, redisLocker *redislock.Client, message amqp.Delivery) {
	log.Info().Caller().Str("messageId", message.MessageId).Bytes("message", message.Body).Msg("Received a message")

	if catchProcessed(redisClient, message) {
		return
	}

	// We obtain lock for the message, it can be routed to different category queues and we need to be sure that
	// only one consumer process it. If message is locked, we wait possible timeout time and than reject message with
	// requeue it. In case if other consumer timed out, lock will be expired and we will receive this message once again.
	// Another strategy - use retry strategy for obtain lock, but we need to be sure that we check after obtain that
	// message is still not processed. Better just to requeue message.
	lock, err := obtainLock(message.MessageId, redisLocker)
	if err != nil {
		log.Info().Caller().Err(err).Str("messageId", message.MessageId).
			Msg("Message already locked, sleep and requeue processing")

		catchLocked(message)

		return
	}

	defer releaseLock(lock)

	log.Info().Caller().Str("messageId", message.MessageId).Msg("Process message")

	timeout := false // for simulate timeout - set to true

	if timeout {
		// In case if processing is not done within proper time, we sent message back again to queue.
		// In case if consumer dies, RabbitMQ will do it for us https://www.rabbitmq.com/tutorials/tutorial-two-python.html:
		//    "There aren't any message timeouts; RabbitMQ will redeliver the message when the consumer dies.
		//       It's fine even if processing a message takes a very, very long time."
		log.Info().Caller().Str("messageId", message.MessageId).Msg("Message timed out")
		time.Sleep(MessageTTL) // simulate timeout

		err := message.Reject(true)
		if err != nil {
			log.Info().Caller().Err(err).Str("messageId", message.MessageId).
				Msg("Can't reject messages")
		}
	} else {
		time.Sleep(MessageTTL) // simulation of consuming process

		// Mark message as processed. To be sure that others will not process it once again.
		// Expiration chosen big enough to be sure that all queues acknowledge message. In general should be at least
		// the same as RabbitMQ queue TTL.
		redisClient.Set(context.Background(), message.MessageId, Processed, 31*24*time.Hour)

		err := message.Ack(false)
		if err != nil {
			log.Info().Caller().Err(err).Str("messageId", message.MessageId).
				Msg("Can't ack messages")
		}

		log.Info().Caller().Str("messageId", message.MessageId).Msg("Message successfully processed")
	}
}

func catchLocked(message amqp.Delivery) {
	time.Sleep(MessageTTL)

	err := message.Reject(true)
	if err != nil {
		log.Info().Caller().Err(err).Str("messageId", message.MessageId).
			Msg("Can't reject messages")
	}
}

func releaseLock(lock *redislock.Lock) {
	err := lock.Release(context.Background())
	if err != nil {
		log.Error().Caller().Err(err).Msg("Failed to release lock record in redis")
	}
}

func catchProcessed(redisClient *redis.Client, message amqp.Delivery) bool {
	v := redisClient.Get(context.Background(), message.MessageId)

	if v.Val() != Processed {
		return false
	}

	// prevent case when message delivered to several queues, but consuming of some queues was not done quick enough
	log.Info().Caller().Str("messageId", message.MessageId).Msg("Message already processed")

	err := message.Ack(false)
	if err != nil {
		log.Info().Caller().Err(err).Str("messageId", message.MessageId).
			Msg("Can't ack messages")
	}

	return true
}

func obtainLock(id string, redisLocker *redislock.Client) (*redislock.Lock, error) {
	lock, err := redisLocker.Obtain(context.Background(), id, MessageTTL, nil)
	if errors.Is(err, redislock.ErrNotObtained) {
		return nil, fmt.Errorf("can't obtain redis lock: %w", err)
	}

	if err != nil {
		log.Error().Caller().Err(err).Msg("Failed to lock record in redis")

		os.Exit(1)
	}

	return lock, nil
}

This repository is proof of concept for the following task:

We have stream of objects that should be moderated. Objects can belongs to several categories.
Moderators can subscribe to several categories and we should show object for moderation only to on moderator.
Moderation process should support timeouts, when moderator lost connection, or similar.

RabbitMQ used for delivery messages to consumers.
Redis used for store process state and as lock manager.

Run "docker-compose up" to start Redis and RabbitMQ

To setup RabbitMQ queues and exchanges run:
go get github.com/streadway/amqp
go get github.com/bsm/redislock
go get github.com/go-redis/redis
go run init.go

To receive message run in separate tabs:
go run receive.go -category category1
go run receive.go -category category2

To send message run:
go run send.go -id 1

Message will be processed only by one receiver, if message was processed before, 
or it is in processing - it will be skipped. 
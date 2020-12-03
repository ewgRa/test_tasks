format:
	gofumpt -w .
	gci -w .

# Run lint checks, used for CI
lint:
	golangci-lint run --enable-all --exclude-use-default=false --disable exhaustivestruct

test:
	GIN_MODE=release go test -race -v ./...

# Shell into local docker api instance
shell-docker-api:
	docker exec -w /app -ti system-design-moderation-queue-api /bin/bash

start-consumer:
	docker exec -w /app/cmd/consumer -ti system-design-moderation-queue-api bash -c 'go run *.go -topic ${topic} -categories ${categories}'

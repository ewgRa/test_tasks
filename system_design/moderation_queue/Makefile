up:
	docker-compose up

# Shell into local docker api instance
shell-docker-api:
	docker-compose exec -w /app api /bin/bash

start-consumer:
	docker-compose exec -w /app/cmd/consumer api bash -c 'go run *.go -topic ${topic} -categories ${categories}'

format:
	gofumpt -w .
	gci -w .

# Run lint checks, used for CI
lint:
	golangci-lint run --enable-all --exclude-use-default=false --disable exhaustivestruct

test:
	GIN_MODE=release go test -race -v ./...

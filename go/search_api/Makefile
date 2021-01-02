.PHONY: build test

# Build production ready docker image
build:
	docker build -t test-task-go-search-api -f build/Dockerfile .

up:
	docker-compose up

# Format code
format:
	gofumpt -w .
	gci -w .

# Run lint checks, used for CI
lint:
	golangci-lint run --enable-all --exclude-use-default=false --disable exhaustivestruct --exclude "Range statement for test .* does not use range value in test Run"

# Shell into local docker api instance
shell-docker-api:
	docker-compose exec -w /app api /bin/bash

# Run tests, run it inside local docker api instance (make docker-api -> make test)
test:ES_INDEX=products_test
test:
	$(MAKE) load-test-data
	GIN_MODE=release go test -race -v ./...

# Load data to elasticsearch test index, used for performing integration tests
load-test-data:
	ES_INDEX=products_test $(MAKE) load-data

# Load data to elasticsearch local index, usually used once for fill elasticsearch with example data
load-local-data:
	ES_INDEX=products_local $(MAKE) load-data

# internal Makefile target
load-data:ES_INDEX_URL=${ES_URL}/${ES_INDEX}
load-data:CURL=curl --silent --output /dev/null --show-error --fail
load-data:
	${CURL} -XDELETE ${ES_INDEX_URL}?ignore_unavailable=true
	${CURL} -XPUT ${ES_INDEX_URL} -H "Content-Type: application/json" --data-binary @build/es.mapping.json
	${CURL} -XPOST ${ES_INDEX_URL}/_bulk -H "Content-Type: application/x-ndjson" --data-binary @test/es.data.bulk
	${CURL} ${ES_INDEX_URL}/_refresh

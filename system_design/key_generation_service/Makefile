.PHONY: up
up:
	docker-compose up --scale api=3

# Shell into local docker api instance
.PHONY: shell-docker-api
shell-docker-api:
	docker-compose exec api /bin/bash

.PHONY: test
test:
	mvn test -Dspring.profiles.active=test

.PHONY: lint
lint:
	mvn formatter:validate

.PHONY: format
format:
	mvn formatter:format

.PHONY: build
build:
	docker build -t test-task-system-design-key-generation-service -f build/Dockerfile .

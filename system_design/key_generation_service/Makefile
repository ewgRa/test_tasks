up:
	docker-compose up --scale api=3

# Shell into local docker api instance
shell-docker-api:
	docker-compose exec api /bin/bash

test:
	mvn test

lint:
	mvn formatter:validate

format:
	mvn formatter:format
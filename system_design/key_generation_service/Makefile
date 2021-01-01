up:
	docker-compose up --scale api=3

test:
	mvn test

# Shell into local docker api instance
shell-docker-api:
	docker-compose exec api /bin/bash

format:
	mvn formatter:format
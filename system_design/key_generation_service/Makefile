up:
	docker-compose up --scale api=3

up-stress-test:
	docker-compose -f build/docker-compose-stress-test.yml up
# docker-compose run -u "$(id -u):$(id -g)" k6 run --logformat raw --console-output=/scripts/test.csv /scripts/stages.js

test:
	mvn test

# Shell into local docker api instance
shell-docker-api:
	docker exec -w /app -ti system-design-key-generation-service-api /bin/bash

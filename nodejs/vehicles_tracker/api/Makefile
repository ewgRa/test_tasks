.PHONY: build test

# Shell into local docker api instance
docker-api:
	docker exec -w /app -ti nodejs-test-task-api /bin/sh

docker-kafka:
	docker exec -ti nodejs-test-task-kafka /bin/bash

# Run tests, run it inside local docker api instance (make docker-api -> make test)
test:ES_INDEX=shuttle_allygator_test
test:KAFKA_TOPIC=test
test:
	$(MAKE) load-test-data
	npm test

lint:
	./node_modules/.bin/eslint src/**

# Load data to elasticsearch test index, used for performing integration tests
load-test-data:
	ES_INDEX=shuttle_allygator_test $(MAKE) load-data

# Load data to elasticsearch local index, usually used once for fill elasticsearch with example data
load-local-data:
	ES_INDEX=shuttle_allygator_local $(MAKE) load-data

# internal Makefile target
load-data:ES_INDEX_URL=${ES_URL}/${ES_INDEX}
load-data:CURL=curl --silent --output /dev/null --show-error --fail
load-data:
	${CURL} -XDELETE ${ES_INDEX_URL}?ignore_unavailable=true
	${CURL} -XPUT ${ES_INDEX_URL} -H "Content-Type: application/json" --data-binary @build/es.mapping.json
	${CURL} ${ES_INDEX_URL}/_refresh


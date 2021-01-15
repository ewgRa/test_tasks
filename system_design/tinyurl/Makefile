.PHONY: up
up:
	docker-compose up

.PHONY: init
init:
	docker-compose up -d postgres
	docker-compose exec -u postgres postgres createdb tiny_url
	docker-compose up -d shorten-api
	docker-compose exec shorten-api mvn flyway:migrate
	timeout 120 bash -c 'until docker-compose ps shorten-api | grep -q \(healthy\); do sleep 2; done' || (docker-compose ps && false)
	docker-compose exec shorten-api curl -XPUT key-generation-service:8080/counter-based/init
	docker-compose stop

.PHONY: test
test:
	docker-compose exec shorten-api make test
	docker-compose exec visitor-app make test

.PHONY: lint
lint:
	cd shorten-api && make lint
	cd visitor-app && make lint

.PHONY: format
format:
	cd shorten-api && make format
	cd visitor-app && make format

.PHONY: up
up:
	docker-compose up

.PHONY: init
init:
	docker-compose exec shorten-api curl -XPUT key-generation-service:8080/counter-based/init
	docker-compose exec -u postgres postgres createdb tiny_url
	docker-compose exec shorten-api mvn flyway:migrate

.PHONY: test
test:
	cd shorten-api && make test
	cd visitor-app && make test

.PHONY: lint
lint:
	cd shorten-api && make lint
	cd visitor-app && make lint

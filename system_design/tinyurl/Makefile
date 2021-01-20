.PHONY: up
up:
	docker-compose up

.PHONY: init
init:
	docker-compose up -d postgres
	timeout 60 bash -c 'until docker-compose exec -T -u postgres postgres psql -c "select 1"; do sleep 2; done' || (docker-compose ps && false)
	docker-compose exec -T -u postgres postgres createdb tiny_url_dev
	docker-compose exec -T -u postgres postgres createdb tiny_url_test
	docker-compose up -d shorten-api
	docker-compose exec -T shorten-api mvn flyway:migrate -Dflyway.user=postgres -Dflyway.password=G3sfF7@tFd# -Dflyway.url=jdbc:postgresql://postgres:5432/tiny_url_dev
	docker-compose exec -T shorten-api mvn flyway:migrate -Dflyway.user=postgres -Dflyway.password=G3sfF7@tFd# -Dflyway.url=jdbc:postgresql://postgres:5432/tiny_url_test
	timeout 60 bash -c 'until docker-compose ps shorten-api | grep -q \(healthy\); do sleep 2; done' || (docker-compose ps && false)
	docker-compose exec -T shorten-api curl -XPUT key-generation-service:8080/counter-based/init
	timeout 60 bash -c 'until docker-compose ps key-generation-service-test | grep -q \(healthy\); do sleep 2; done' || (docker-compose ps && false)
	docker-compose exec -T shorten-api curl -XPUT key-generation-service-test:8080/counter-based/init
	docker-compose stop

.PHONY: test
test:
	docker-compose exec -T shorten-api make test
	docker-compose exec -T visitor-app make test

.PHONY: lint
lint:
	cd shorten-api && make lint
	cd visitor-app && make lint

.PHONY: format
format:
	cd shorten-api && make format
	cd visitor-app && make format

init:
	docker-compose exec shorten-api curl -XPUT key-generation-service:8080/counter-based/init
	docker-compose exec -u postgres postgres createdb tiny_url
	docker-compose exec shorten-api mvn flyway:migrate

init:
	docker-compose exec user-api curl -XPUT user-api-key-generation-service:8080/counter-based/init
	docker-compose exec -u postgres postgres createdb tiny_url
	docker-compose exec user-api mvn flyway:migrate

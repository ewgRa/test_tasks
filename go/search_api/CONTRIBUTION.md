# Setup a local development environment

- Run `sh build/dev.init.sh` to check requirements and initial setup
- Read output and follow instructions in case of needs
- Run `docker-compose up` to start docker-compose
- Open a new terminal
- Run `make shell-docker-api` to shell into docker instance
- Run `make load-local-data` to fill Elasticsearch with some products

Now you can open swagger and play with API: http://localhost:8090

# Development

API instance can be accessed on http://localhost:8080.
Swagger linked to this API can be accessed on http://localhost:8090.
For more information, you can use docker-compose.yml as documentation.

In Makefile, you will find a set of useful commands that make it easier to develop.

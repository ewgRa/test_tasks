# Setup local environment

1. run: **sh build/dev.init.sh** to check requirements and initial setup
2. read output and follow instructions in case of needs
3. run: **docker-compose up** for start docker-compose
4. **open new terminal**
5.  run: **make docker-api** to shell into docker instance
6. run: **make load-local-data** to fill elasticsearch with some products

Now you can open swagger and play with API: http://localhost:8090

# Development

API instance can be accessed on http://localhost:8080.
Swagger linked to this API can be accessed on http://localhost:8090.
For more information you can use docker-compose.yml as documentation.

In Makefile you will find set of useful commands that make it easier to develop. 
For example, to run tests start docker-compose, than make "docker-api" target and inside docker
instance "test" target.

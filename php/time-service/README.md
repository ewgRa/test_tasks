# Setup with docker

Install [docker-compose](https://docs.docker.com/compose/install/) and run
`docker-compose up`

Open another terminal and run `make install` to install dependencies.

Two services will be up and running:
* Time service, that provide required functionality, http://localhost:8001/. User is "test", password is "test";

* Swagger-UI that can be used as documentation and testing interface: http://localhost:8091/
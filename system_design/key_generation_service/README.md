// FIXME XXX: tests
// FIXME XXX: CI
# Task definition
Develop Key Generation Service that generates unique keys with specific length.
Can be used as part of solution for TinyURL and other services.

# High-level design
## Counter-based
![](docs/high-level-counter-based-design.png)

# Usage
- Run `make up`
- Open in browser Swagger UI http://localhost:8090/
- Initialize counter-based approach http://localhost:8090/#/counter-based/put_counter_based_init
- Generate counter-based keys via http://localhost:8090/#/counter-based/get_counter_based_key

# Development
Please check Makefile and docker-compose.yml for understanding how the solution works under the hood.
#!/bin/sh

docker pull elasticsearch:7.3.1
docker network create esnet-oss;
docker run \
  --rm \
  --publish 9200:9200 \
  --env "node.attr.testattr=test" \
  --env "path.repo=/tmp" \
  --env "repositories.url.allowed_urls=http://snapshot.*" \
  --env "discovery.type=single-node" \
  --network=esnet-oss \
  --name=elasticsearch-oss \
  --detach \
  elasticsearch:7.3.1
docker run --network esnet-oss --rm appropriate/curl --max-time 120 --retry 120 --retry-delay 1 --retry-connrefused --show-error --silent http://elasticsearch-oss:9200
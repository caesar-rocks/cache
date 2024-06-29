#!/bin/bash

# exposes default port 6379 to local host
docker  run --rm -it -p 6379:6379 --name some-redis -d redis

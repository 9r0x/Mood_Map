#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..

cd $ROOT

docker build -f ./Dockerfile_postgres -t gaocegege/postgres:9 .
docker push gaocegege/postgres:9

cd - > /dev/null

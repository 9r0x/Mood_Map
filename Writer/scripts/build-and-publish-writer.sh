#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE}")/..

cd $ROOT

docker run --rm \
       -v `pwd`:/go/src/github.com/gaocegege/hackys-backend-writer \
       -e GOPATH=/go:/go/src/github.com/gaocegege/hackys-backend-writer/vendor golang:1.6 sh \
       -c "cd /go/src/github.com/gaocegege/hackys-backend-writer && go build -o app"
docker build -t gaocegege/hackys-writer .
docker push gaocegege/hackys-writer
cd - > /dev/null

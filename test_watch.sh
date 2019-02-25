#!/bin/bash
while true; do

inotifywait -e modify,create,delete -r ./cmd && \
# wait to build
sleep 5 
## run unit tests
go test -v ./cmd/server/methods/

## trigger integration tests
npm run testCI --prefix integration-tests/

done
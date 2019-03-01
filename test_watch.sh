#!/bin/bash
while true; do

inotifywait -e modify,create,delete -r ./cmd && \
clear
# wait to build
# sleep 5 
## run unit and integration tests
go test -v ./cmd/server/methods/ -covermode=count -coverprofile=./cmd/server/methods/coverage.out && \
npm run testCI --prefix integration-tests/

done
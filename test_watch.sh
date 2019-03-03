#!/bin/bash
while true; do

run_tests() {
	go test -v ./cmd/server/methods/ -covermode=count -coverprofile=./cmd/server/methods/coverage.out
	if [ $? -eq 0 ]; then
		npm run testCI --prefix integration-tests/
	fi
}

inotifywait -e modify,create,delete -r ./cmd && \
	clear
	run_tests
done
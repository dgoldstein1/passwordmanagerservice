#!/bin/bash
while true; do

run_tests() {
	go test -v ./cmd/server/methods/ -covermode=count
	if [ $? -eq 0 ]; then
		sleep 5
		npm run testCI --prefix integration-tests/
	fi
}

inotifywait -e modify,create,delete -r ./cmd && \
	clear
	run_tests
done
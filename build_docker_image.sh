#!/bin/bash

./create_documentation.sh &> /dev/null
docker build -t dgoldstein1/passwordservice:latest .
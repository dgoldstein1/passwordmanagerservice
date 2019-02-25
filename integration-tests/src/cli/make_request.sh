#!/bin/bash

# make_request.sh

CA='./src/cli/certs/client.trust.pem'
CERT='./src/cli/certs/client.cert.pem'
KEY='./src/cli/certs/client.key.pem'
USE_TLS=false

# parse command line args
URL=$1
METHOD=$2
DATA=$3

if [ "$USE_TLS" == "false" ]; then
	if [ "$METHOD" == "GET" ]; then
		curl ${URL}
	elif [ "$METHOD" == "DEL" ]; then
	   curl \
		   	--header "Content-Type: application/json" \
		   	--request "DELETE" \
		   	--data ''"$DATA"'' ${URL}
	else
		curl \
			--header "Content-Type: application/json" \
			--request $METHOD \
			--data ''"$DATA"'' ${URL}
	fi
else
	if [ "$METHOD" == "GET" ]; then
		curl --insecure \
			--cacert ${CA} \
			--cert ${CERT} \
			--key ${KEY} ${URL}
    elif [ "$METHOD" == "DEL" ]; then
        curl \
        	--header "Content-Type: application/json" \
        	--request "DELETE" \
        	--data ''"$DATA"'' \
        	--insecure \
        	--cacert ${CA} \
        	--cert ${CERT} \
        	--key ${KEY} ${URL}
	else
		curl \
			--header "Content-Type: application/json" \
			--request $METHOD \
			--data ''"$DATA"'' \
			--insecure \
			--cacert ${CA} \
			--cert ${CERT} \
			--key ${KEY} ${URL}
	fi
fi


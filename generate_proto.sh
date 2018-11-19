#!/bin/bash

# logs success / failure of last command
# param name of command
log_success_or_failure() {
    if [ $? -eq 0 ]
    then
        echo "$(tput setab 2 )--- SUCCESS --- ${1} $(tput sgr0)"
    else
        echo "$(tput setab 1 )--- FAILURE --- ${1} $(tput sgr0)"
       	exit $?
    fi
}

cd ..

serviceDir=passwordservice/cmd/server/
tempDir=passwordservice/tempDir

# copy things over to temp dir
if [ -d $serviceDir ]; then
	mkdir $tempDir
	cp "${serviceDir}/methods/new_server.go" "${tempDir}/new_server.go"
	cp "${serviceDir}/gateway_proxy.go" "${tempDir}/gateway_proxy.go" 
	log_success_or_failure "Copying temp new_server"
fi

fabric --generate "passwordservice"
log_success_or_failure "Generating space service proto "
echo "$(tput setab 2 ) Proto generated succesfully  $(tput sgr0) ${1} "

# copy back temp and cleanup
if [ -d $tempDir ]; then
	cp "${tempDir}/new_server.go" "${serviceDir}/methods/new_server.go" 
	log_success_or_failure "Restoring new_server.go"
	cp "${tempDir}/gateway_proxy.go" "${serviceDir}/gateway_proxy.go"  
	log_success_or_failure "Restoring gateway_proxy.go"
	rm -rf $tempDir
	log_success_or_failure "Removing tmp dir"
fi


cd passwordservice

# add in "bson" attribute into protofile
sed -i'' -e  's/`json:"-"`/`json:"-" bson:"-"`/g' ./protobuf/passwordservice.pb.go
log_success_or_failure "Added 'bson' to XXX_ fields"

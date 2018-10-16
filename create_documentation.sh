#!/bin/bash

# createDocumentation.sh

# use ci assumes you don't have aglio installed on your machine
USE_CI=$1
documentationDir="./docs/api"
if [ "$USE_CI" = "true" ]; then
    # make temp package.json
    > package.json
    echo '{
      "name": "spaceservice",
      "version": "1.0.0",
      "description": "A RESTful service written in Go for managing spaces.",
      "main": "index.js",
      "directories": {
        "doc": "docs",
        "test": "test"
      },
      "scripts": {
        "aglio" : "aglio"
      },
      "repository": {
        "type": "git",
        "url": "ssh://git@cdn-gitlab.363-283.io:2252/chimera/spaceservice.git"
      },
      "author": "",
      "license": "ISC",
      "dependencies": {
        "aglio": "^2.3.0"
      }
    }
    ' >> package.json
    # install
    npm install
    # create aglio docs
    npm run aglio -- -i ${documentationDir}/index.apib -o ${documentationDir}/index.html
    npm run aglio -- -i ${documentationDir}/rest.apib -o ${documentationDir}/rest.html
    npm run aglio -- -i ${documentationDir}/changelog.apib -o ${documentationDir}/changelog.html
    npm run aglio -- -i ${documentationDir}/datamodel.apib -o ${documentationDir}/datamodel.html
    # remove temp dirs
    rm -rf node_modules
    rm package.json

else
	# create aglio docs
	aglio -i ${documentationDir}/index.apib -o ${documentationDir}/index.html
	aglio -i ${documentationDir}/rest.apib -o ${documentationDir}/rest.html
	aglio -i ${documentationDir}/changelog.apib -o ${documentationDir}/changelog.html
	aglio -i ${documentationDir}/datamodel.apib -o ${documentationDir}/datamodel.html
fi



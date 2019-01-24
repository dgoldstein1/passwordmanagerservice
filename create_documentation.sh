#!/bin/bash

# createDocumentation.sh
documentationDir="./docs/api"

# create aglio docs
aglio -i ${documentationDir}/index.apib -o ${documentationDir}/index.html
aglio -i ${documentationDir}/rest.apib -o ${documentationDir}/rest.html
aglio -i ${documentationDir}/changelog.apib -o ${documentationDir}/changelog.html
aglio -i ${documentationDir}/datamodel.apib -o ${documentationDir}/datamodel.html



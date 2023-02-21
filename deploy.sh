#!/bin/bash -xe
# title           :deploy.sh
# description     :Deploy this code to Lambda
# date            :2021-11-28
# version         :0.1.0
# usage           :bash deploy.sh
# assumption      :The AWS function already exists. Zip program is available, Bash is available
# variables
#                 :BUILD_PATH - Path put the build zip which will be sent to AWS
#                 :FUNC_NAME  - Name of the AWS lambda function to update
#------------------------------------------------------------------------------

rm -f -- main function.zip 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main ./cmd/lambda
zip -r function.zip main templates/

BUILD_PATH=${RELEASE_PATH:-$PWD}
FUNC_NAME=${FUNC_NAME:-generateUUID}
ARCHIVE=${LAMBDA_ARCHIVE:-function.zip}

aws lambda update-function-code \
    --function-name $FUNC_NAME \
    --zip-file fileb://$BUILD_PATH/$ARCHIVE

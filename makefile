#!/bin/bash
export GOPATH="/Users/amarsharma/Dev/website"
cd bin
go build app
cd ../src
../bin/app

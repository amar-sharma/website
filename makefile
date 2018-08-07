#!/bin/bash
export GOPATH="/home/legitderank/website"
mkdir bin
cd bin
/usr/local/go/bin/go build app
cd ../src
../bin/app
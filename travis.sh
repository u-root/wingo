#!/bin/bash
if [ -z "${GOPATH}" ]; then
        export GOPATH=/home/travis/gopath
fi
set -e
echo "-----------------------> generate test"
go generate

echo "-----------------------> Initial build test"
go build

cp wingo wingo.1st
echo "-----------------------> Second build test"
go generate
go build

echo "-----------------------> Reproducible test"
cmp wingo wingo.1st

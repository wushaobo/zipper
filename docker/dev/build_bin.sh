#!/usr/bin/env bash

set -e

app_name="zipper"

cd $(dirname $0)/../..


function export_gopath() {
    export GOPATH=${PWD}
}

function clean() {
    echo "##cleaning binary ..."
    rm -rf bin/*        # generated by  go install
    rm -f ${app_name}     # generated by  go build
}

function dependencies() {
    echo "##install dependencies ... (It requires to cross the GFW if you are in China mainland.)"
    echo "  -> github.com/gorilla/mux"
    go get github.com/gorilla/mux
    echo "  -> gopkg.in/redis.v3"
    go get gopkg.in/redis.v3
}

function build_binary() {
    echo "##building binary ..."
    GOOS=linux GOARCH=amd64 go install ${app_name}
}


echo "Start building bin ..."
export_gopath
clean
dependencies
build_binary

echo "Built bin."

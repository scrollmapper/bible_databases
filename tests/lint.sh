#!/bin/bash
FormatCheck=$(gofmt -l *.go | wc -l)
if [ $FormatCheck -gt 0 ]; then
    gofmt -l *.go
    echo "gofmt -w *.go your code please."
    exit 1
fi
## Run golint
golint -set_exit_status
if [ $? -gt 0 ]; then
    exit 1
fi

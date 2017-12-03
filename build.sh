#!/bin/sh

# Build a statically-linked binary
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

docker build -t nwwa .

#!/usr/bin/env bash

protoc --proto_path=api/proto --go_out=plugins=grpc:pkg/ ./api/proto/example.proto
protoc --proto_path=api/proto \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:pkg/ \
     ./api/proto/example.proto


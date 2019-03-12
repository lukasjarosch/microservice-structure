#!/usr/bin/env bash

# generate GO stubs from proto file
protoc -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=api/proto --go_out=plugins=grpc:pkg/ ./api/proto/example.proto

# generate grpc HTTP gateway from proto file
protoc --proto_path=api/proto \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:pkg/ \
     ./api/proto/example.proto

# generate swagger spec of the generated HTTP gateway
protoc -I/usr/local/include -I. \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true:. \
    ./api/proto/example.proto

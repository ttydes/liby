#!/bin/sh

deps(){
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go get golang.org/x/net/http2
}

proto(){
    for f in services/*/*.proto; do \
	protoc --go_out=plugins=grpc:. $$f; \
	echo compiled: $$f; \
    done
}

certs(){
    ./scripts/mkcerts
}

$1

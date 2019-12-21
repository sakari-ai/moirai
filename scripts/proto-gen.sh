#!/bin/sh

PROTO_FILES=$(find proto -name "*.proto")
for PROTO_FILE in ${PROTO_FILES}; do
    protoc -I/usr/local/include -Iproto \
        -I${GOPATH}/src \
        -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
        -Ithird_party/googleapis \
        --go_out=plugins=grpc:${GOPATH}/src \
        ${PROTO_FILE}
    protoc -I/usr/local/include -Iproto \
        -I${GOPATH}/src \
        -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
        -Ithird_party/googleapis \
        --grpc-gateway_out=logtostderr=true:${GOPATH}/src \
        ${PROTO_FILE}
    protoc -I/usr/local/include -Iproto \
        -I${GOPATH}/src \
        -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
        -Ithird_party/googleapis \
        --swagger_out=logtostderr=true:proto \
        ${PROTO_FILE}
done
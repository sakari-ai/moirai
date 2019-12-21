#!/bin/bash

GRPC_GATEWAY_VERSION="v1.5.1"
PROTOC_GEN_GRPC_GATEWAY_URL="https://github.com/grpc-ecosystem/grpc-gateway/releases/download/${GRPC_GATEWAY_VERSION}/protoc-gen-grpc-gateway-${GRPC_GATEWAY_VERSION}-darwin-x86_64"
PROTOC_GEN_SWAGGER_URL="https://github.com/grpc-ecosystem/grpc-gateway/releases/download/${GRPC_GATEWAY_VERSION}/protoc-gen-swagger-${GRPC_GATEWAY_VERSION}-darwin-x86_64"

verifyOS() {
  if [ "$(uname)" != "Darwin" ]
  then
        echo "Only supported for macOS"
        return 1
  fi
  return 0
}

verifyGOPATH() {
  if [ -z "${GOPATH}" ]
  then
        echo "\${GOPATH} is required"
        return 1
  fi
  return 0
}

prepareGOPATHBin() {
  mkdir -p ${GOPATH}/bin
}

installProtocGenGo() {
  go install ./vendor/github.com/golang/protobuf/protoc-gen-go
}

installGRPCGateway() {
  echo "Downloading protoc-gen-grpc-gateway-${GRPC_GATEWAY_VERSION}..."
  curl -L ${PROTOC_GEN_GRPC_GATEWAY_URL} --out ${GOPATH}/bin/protoc-gen-grpc-gateway

  echo "Downloading protoc-gen-swagger-${GRPC_GATEWAY_VERSION}..."
  curl -L ${PROTOC_GEN_SWAGGER_URL} --out ${GOPATH}/bin/protoc-gen-swagger

  echo "Making file executable {\${GOPATH}/bin/protoc-gen-grpc-gateway, \${GOPATH}/bin/protoc-gen-swagger}"
  chmod +x ${GOPATH}/bin/protoc-gen-grpc-gateway
  chmod +x ${GOPATH}/bin/protoc-gen-swagger
}

bye() {
  echo
  echo "Done"
}

# quit on any error
set -e

echo "Checking \$OS..."
verifyOS

echo "Checking \$GOPATH..."
verifyGOPATH

echo "Preparing \${GOPATH}/bin..."
prepareGOPATHBin

echo "Installing protoc-gen-go..."
installProtocGenGo

echo "Installing grpc-gateway..."
installGRPCGateway

bye








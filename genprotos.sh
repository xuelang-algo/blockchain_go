#!/usr/bin/env bash
protoc services/*.proto --go_out=plugins=gorpc:services -I services/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

protoc services/*.proto  --grpc-gateway_out=logtostderr=true:services -I services/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

protoc services/*.proto  --swagger_out=logtostderr=true:services -I services/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

hostname="blockchain.hz-xuelang.xyz"

osname=`uname`
if [[ ${osname} -eq "Darwin" ]]; then
sed -i "" 's/^  "consumes": \[/  "host": '"\"${hostname}\","'\
  "consumes": \[/' services/*.swagger.json
else
sed -i 's/^  "consumes": \[/  host: '"${hostname}"'\n"consumes": \[/' services/*.swagger.json
fi
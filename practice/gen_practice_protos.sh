#!/usr/bin/env bash
protoc *.proto --go_out=plugins=gorpc:. -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

protoc *.proto  --grpc-gateway_out=logtostderr=true:. -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

protoc *.proto  --swagger_out=logtostderr=true:. -I ./ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

hostname="blockchain.hz-xuelang.xyz"

osname=`uname`
if [[ ${osname} -eq "Darwin" ]]; then
sed -i "" 's/^  "consumes": \[/  "host": '"\"${hostname}\","'\
  "consumes": \[/' ./*.swagger.json
else
sed -i 's/^  "consumes": \[/  host: '"${hostname}"'\n"consumes": \[/' ./*.swagger.json
fi
#! /usr/bin/env bash

set -e

if ! [[ -x "$(command -v protoc-gen-go)" ]]; then
  go get github.com/golang/protobuf/protoc-gen-go
  go install github.com/golang/protobuf/protoc-gen-go

  if ! [[ -x "$(command -v protoc-gen-go)" ]]; then
    echo "Failed to instal protoc-gen-go, check your GOPATH"
    exit 1
  fi
fi


PROTOC_OPTS="-I../third_party/googleapis/ -I../apis/ --go_out=plugins=grpc:temp"

mkdir -p temp/

protoc $PROTOC_OPTS ../apis/tinkoff/cloud/stt/v1/*.proto
protoc $PROTOC_OPTS ../apis/tinkoff/cloud/tts/v1/*.proto
protoc $PROTOC_OPTS ../apis/tinkoff/cloud/longrunning/v1/*.proto


rm -rf pkg/tinkoff

mv temp/github.com/TinkoffCreditSystems/voicekit-examples/golang/pkg/* pkg

rm -rf temp/
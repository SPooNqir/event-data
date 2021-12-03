#!/bin/bash

GEN_PATH="."
GOPATH=$(go env GOPATH)

AUTHORIZATIONS_VERSION="master"
curl -o proto/authorizations.proto "https://raw.githubusercontent.com/slavayssiere-spoon/authorizations/$AUTHORIZATIONS_VERSION/proto/authorizations.proto"

GROUPS_VERSION="master"
curl -o proto/groups.proto "https://raw.githubusercontent.com/slavayssiere-spoon/groups/$GROUPS_VERSION/proto/groups.proto"

USERS_VERSION="master"
curl -o proto/users.proto "https://raw.githubusercontent.com/slavayssiere-spoon/users/$USERS_VERSION/proto/users.proto"

ROBOTS_VERSION="master"
curl -o proto/robots.proto "https://raw.githubusercontent.com/slavayssiere-spoon/robots/$ROBOTS_VERSION/proto/robots.proto"

HEALTH_VERSION="master"
curl -o proto/health.proto "https://raw.githubusercontent.com/slavayssiere-spoon/health/$HEALTH_VERSION/proto/health.proto"

echo "gen bq_field"
protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  proto/bq_field.proto

echo "gen bq_table"
protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  proto/bq_table.proto

echo "gen identityevent"
protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  proto/IdentityEvent.proto

echo "gen identityevent - add tag"
protoc-go-inject-tag -input=./IdentityEvent.pb.go

echo "gen alert overflow"
protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  proto/DataOverflowAlert.proto

echo "gen alert overflow - add tag"
protoc-go-inject-tag -input=./DataOverflowAlert.pb.go

echo "gen alert not connected"
protoc \
  -I proto \
  -I $GOPATH/src/include \
  --go_out=$GEN_PATH      --go_opt=paths=source_relative \
  proto/RobotNotConnectedAlert.proto

echo "gen alert not connected - add tag"
protoc-go-inject-tag -input=./RobotNotConnectedAlert.pb.go

echo "mod tidy"
go mod tidy

echo "update"
go get -u ./...


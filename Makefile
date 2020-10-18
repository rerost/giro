BIN_DIR := ${PWD}/bin
export PATH := ${BIN_DIR}:${PATH}

PHONY: setup
setup:
	GO111MODULE=off go get github.com/izumin5210/gex/cmd/gex
	gex --build

PHONY: protoc
protoc: 
	protoc --go_out=${GOPATH}/src protos/hosts.proto

PHONY: generate
generate: setup
	go mod tidy
	go generate ./...

.PHONY: mock
mock:
	mockgen github.com/rerost/giro/domain/grpcreflectiface Client > mock/grpcreflectiface/client_test.go

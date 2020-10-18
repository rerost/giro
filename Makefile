BIN_DIR := ${PWD}/bin
export PATH := ${BIN_DIR}:${PATH}

PHONY: setup
setup:
	GO111MODULE=off go get github.com/izumin5210/gex/cmd/gex
	gex --build

PHONY: protoc
protoc: 
	protoc --go_out=${GOPATH}/src e2etest/dummyserver/echo.proto

PHONY: generate
generate: setup
	go mod tidy
	go generate ./...

PHONY: build
build: setup generate
	go build -o ${BIN_DIR}/giro ./cmd/giro
	go build -o ${BIN_DIR}/protoc-gen-reflection-server ./cmd/protoc-gen-reflection-server

.PHONY: mock
mock:
	mockgen github.com/rerost/giro/domain/grpcreflectiface Client > mock/grpcreflectiface/client_test.go

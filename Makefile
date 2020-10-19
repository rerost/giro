BIN_DIR := ${PWD}/bin
export PATH := ${BIN_DIR}:${PATH}

PHONY: setup
setup:
	GO111MODULE=off go get github.com/izumin5210/gex/cmd/gex
	gex --build

PHONY: testcase
testcase:
	protoc --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/onefile.pb runner/genreflectionserver/testprotos/onefile/*.proto
	protoc --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/multifile.pb runner/genreflectionserver/testprotos/multifile/*.proto
	protoc -I=/usr/local/include/ -I=. --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/with_host_option.pb runner/genreflectionserver/testprotos/with_host_option/*.proto

PHONY: protoc
protoc: 
	protoc --go_out=plugins=grpc,paths=source_relative:. e2etest/dummyserver/echo.proto
	protoc -I=/usr/local/include/ -I=. --go_out=plugins=grpc:${GOPATH}/src protos/hosts.proto

PHONY: generate
generate: setup testcase protoc
	go mod tidy
	go generate ./...

PHONY: build
build: setup generate
	go build -ldflags="-X main.Version=0.0.0 -X main.Revision=testhash" -o ${BIN_DIR}/giro ./cmd/giro
	go build -ldflags="-X main.Version=0.0.0 -X main.Revision=testhash" -o ${BIN_DIR}/protoc-gen-reflection-server ./cmd/protoc-gen-reflection-server

PHONY: test
test:
	go test -v ./...

.PHONY: mock
mock:
	mockgen github.com/rerost/giro/domain/grpcreflectiface Client > mock/grpcreflectiface/client_test.go

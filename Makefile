BIN_DIR := ${PWD}/bin
export PATH := ${BIN_DIR}:${PATH}

PHONY: setup
setup:
	go mod download
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

PHONY: testcase
testcase:
	protoc --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/onefile.pb runner/genreflectionserver/testprotos/onefile/*.proto
	protoc --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/multifile.pb runner/genreflectionserver/testprotos/multifile/*.proto
	protoc -I=/usr/local/include/ -I=. --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/with_host_option.pb runner/genreflectionserver/testprotos/with_host_option/*.proto

PHONY: protoc
protoc: 
	protoc -I=/usr/local/include/ -I=. --go-grpc_out=plugins=grpc,paths=source_relative:. e2etest/dummyserver/echo.proto
	protoc -I=/usr/local/include/ -I=. --go-grpc_out=plugins=grpc:${GOPATH}/src rerost/giro/hosts.proto

PHONY: generate
generate: setup testcase protoc
	go generate ./...

PHONY: build
build: setup generate
	go build -ldflags="-X main.Version=0.0.0 -X main.Revision=testhash" -o ${BIN_DIR}/giro ./cmd/giro
	go build -ldflags="-X main.Version=0.0.0 -X main.Revision=testhash" -o ${BIN_DIR}/protoc-gen-reflection-server ./cmd/protoc-gen-reflection-server

PHONY: test
test:
	go test -v ./...

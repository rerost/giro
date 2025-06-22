BIN_DIR := ${PWD}/bin
export PATH := ${BIN_DIR}:${PATH}

PHONY: setup
setup:
	go mod download

PHONY: testcase
testcase:
	protoc --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/onefile.pb runner/genreflectionserver/testprotos/onefile/*.proto
	protoc --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/multifile.pb runner/genreflectionserver/testprotos/multifile/*.proto
	protoc -I=/usr/local/include/ -I=. --include_imports --include_source_info --descriptor_set_out=runner/genreflectionserver/testdata/with_host_option.pb runner/genreflectionserver/testprotos/with_host_option/*.proto

PHONY: protoc
protoc: 
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

	protoc -I=/usr/local/include/ -I=. --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative e2etest/dummyserver/echo.proto
	protoc -I=/usr/local/include/ -I=. --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative rerost/giro/hosts.proto

PHONY: generate
generate: setup testcase protoc
	go generate ./...
	go mod tidy

PHONY: build
build: setup generate
	go build -ldflags="-X main.Version=0.0.0 -X main.Revision=testhash" -o ${BIN_DIR}/giro ./cmd/giro
	go build -ldflags="-X main.Version=0.0.0 -X main.Revision=testhash" -o ${BIN_DIR}/protoc-gen-reflection-server ./cmd/protoc-gen-reflection-server

PHONY: test
test:
	go test -v ./...

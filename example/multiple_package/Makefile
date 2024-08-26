.PHONY: setup
setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

	go install github.com/rerost/giro/cmd/protoc-gen-reflection-server
	go install github.com/rerost/giro/cmd/giro

.PHONY: protoc
protoc: setup
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --reflection-server_out=. $(shell find  . -name '*.proto')
	grpc_tools_ruby_protoc --ruby_out=$(PWD)/server --grpc_out=$(PWD)/server $(shell find  . -name '*.proto')

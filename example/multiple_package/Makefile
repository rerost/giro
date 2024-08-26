
.PHONY: protoc
protoc:
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative --reflection-server_out=. $(shell find  . -name '*.proto')
	grpc_tools_ruby_protoc --ruby_out=$(PWD)/server --grpc_out=$(PWD)/server $(shell find  . -name '*.proto')


.PHONY: protoc
protoc:
	protoc --go_out=plugins=grpc,paths=source_relative:.  --reflection-server_out=. $(shell find  . -name '*.proto')
	grpc_tools_ruby_protoc --ruby_out=$(PWD)/server --grpc_out=$(PWD)/server $(shell find  . -name '*.proto')

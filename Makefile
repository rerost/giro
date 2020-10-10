PHONY: protoc
protoc: 
	protoc --go_out=${GOPATH}/src protos/hosts.proto

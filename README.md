# giro
[![Go Report Card](https://goreportcard.com/badge/github.com/rerost/giro)](https://goreportcard.com/report/github.com/rerost/giro)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/rerost/giro)](http://github.com/rerost/giro/releases/latest)
[![license](https://img.shields.io/github/license/rerost/giro.svg)](./LICENSE)

An alternative to [`grpc_cli`](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md).
giro can be used for gRPC servers without Server Reflection

## Installation
### macOS
```
brew install rerost/tools/giro
```

### Linux
https://github.com/rerost/giro/releases


## Tutorial
```
$ git clone https://github.com/rerost/giro.git giro
$ cd giro/example/multiple_package
```

### Create & Run reflection server
```
$ protoc --go_out=plugins=grpc,paths=source_relative:.  --reflection-server_out=. $(find  . -name '*.proto')
$ go run main.go
```

### Run gRPC Server
```
cd example/multiple_package/server
docker build -t test .
docker run -it -p 5001:5001 test
cd ../../../
```

### Unary call with giro
```
$ giro ls
example.multiple_package.protos.one.GiroService
example.multiple_package.protos.twofile.BqvService
grpc.health.v1.Health
grpc.reflection.v1alpha.ServerReflection
rerost.giro.v1.HostService
$ giro call --rpc-server=localhost:5001 example.multiple_package.protos.one.GiroService/GiroTest1 '{"message": "test"}'
{"message":"test"}
```

# giro
[![Go Report Card](https://goreportcard.com/badge/github.com/rerost/giro)](https://goreportcard.com/report/github.com/rerost/giro)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/rerost/giro)](http://github.com/rerost/giro/releases/latest)
[![license](https://img.shields.io/github/license/rerost/giro.svg)](./LICENSE)
[![codecov](https://codecov.io/gh/rerost/giro/graph/badge.svg?token=1C3JHYSTTB)](https://codecov.io/gh/rerost/giro)

An alternative to [`grpc_cli`](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md).
giro can be used for gRPC servers without Server Reflection

## Installation
```
brew install rerost/tools/giro
```

### Create & Run reflection server
```
$ protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --reflection-server_out=. $(shell find  . -name '*.proto')
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

## Develop
### Update snapshot
```
$ go test e2etest/giro_test.go -update
```

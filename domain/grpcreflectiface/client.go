package grpcreflectiface

import (
	"sync"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

type Client interface {
	ListServices() ([]string, error)
	ResolveService(serviceName string) (protoreflect.ServiceDescriptor, error)
	ResolveMessage(messageName string) (proto.Message, error)
}

type Stream = grpc_reflection_v1.ServerReflection_ServerReflectionInfoClient

type clientImpl struct {
	stream Stream

	streamLock sync.Mutex
}

func NewClient(stream Stream) Client {
	return &clientImpl{
		stream: stream,
	}
}

func (c *clientImpl) ListServices() ([]string, error) {
	services := []string{}
	res, err := c.call(&grpc_reflection_v1.ServerReflectionRequest{
		MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_ListServices{
			ListServices: "*",
		},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, service := range res.GetListServicesResponse().Service {
		services = append(services, service.GetName())
	}

	return services, nil
}
func (c *clientImpl) ResolveService(serviceName string) (protoreflect.ServiceDescriptor, error) {
	req := grpc_reflection_v1.ServerReflectionRequest{
		MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_FileContainingSymbol{
			FileContainingSymbol: serviceName,
		},
	}

	resp, err := c.call(&req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var svcDescriptor protoreflect.ServiceDescriptor
	if resp.GetFileDescriptorResponse() == nil {
		return nil, nil
	}

	// ファイルをレジストリに登録
	registry := &protoregistry.Files{}
	checkedFdNames := map[string]bool{}
	registerFunc := func(fdProto *descriptorpb.FileDescriptorProto) error {
		// 重複排除
		if _, ok := checkedFdNames[fdProto.GetName()]; ok {
			return nil
		}
		checkedFdNames[fdProto.GetName()] = true

		file, err := protodesc.NewFile(fdProto, registry)
		if err != nil {
			return errors.WithStack(err)
		}
		return errors.WithStack(registry.RegisterFile(file))
	}

	for _, fdProtoBytes := range resp.GetFileDescriptorResponse().FileDescriptorProto {
		var fdProto descriptorpb.FileDescriptorProto
		if err := proto.Unmarshal(fdProtoBytes, &fdProto); err != nil {
			return nil, errors.WithStack(err)
		}

		if err := c.resolveFiles(registerFunc, &fdProto); err != nil {
			return nil, err
		}
	}
	for _, fdProtoBytes := range resp.GetFileDescriptorResponse().FileDescriptorProto {
		var fdProto descriptorpb.FileDescriptorProto
		if err := proto.Unmarshal(fdProtoBytes, &fdProto); err != nil {
			return nil, errors.WithStack(err)
		}
		file, err := protodesc.NewFile(&fdProto, registry)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		svcDescriptor = file.Services().ByName(protoreflect.FullName(serviceName).Name())
		if svcDescriptor != nil {
			return svcDescriptor, nil
		}
	}
	return nil, errors.New("service not found")
}

func (c *clientImpl) ResolveMessage(messageName string) (proto.Message, error) {
	req := grpc_reflection_v1.ServerReflectionRequest{
		MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_FileContainingSymbol{
			FileContainingSymbol: messageName,
		},
	}

	resp, err := c.call(&req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	registry := &protoregistry.Files{}
	checkedFdNames := map[string]bool{}
	registerFunc := func(fdProto *descriptorpb.FileDescriptorProto) error {
		// 重複排除
		if _, ok := checkedFdNames[fdProto.GetName()]; ok {
			return nil
		}
		checkedFdNames[fdProto.GetName()] = true
		file, err := protodesc.NewFile(fdProto, registry)
		if err != nil {
			return errors.WithStack(err)
		}
		return registry.RegisterFile(file)
	}
	var dynamicMessage proto.Message
	files := map[string]*descriptorpb.FileDescriptorProto{}
	for _, fdProtoBytes := range resp.GetFileDescriptorResponse().FileDescriptorProto {
		var fdProto descriptorpb.FileDescriptorProto

		if err := proto.Unmarshal(fdProtoBytes, &fdProto); err != nil {
			return nil, errors.WithStack(err)
		}
		files[fdProto.GetName()] = &fdProto

		if err := c.resolveFiles(registerFunc, &fdProto); err != nil {
			return nil, err
		}
	}

	// メッセージディスクリプタを探す
	for _, fdProto := range files {
		file, err := protodesc.NewFile(fdProto, registry)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		msgDescriptor := file.Messages().ByName(protoreflect.FullName(messageName).Name())
		if msgDescriptor != nil {
			dynamicMessage = dynamicpb.NewMessage(msgDescriptor)
			break
		}
	}

	return dynamicMessage, nil
}

func (c *clientImpl) call(request *grpc_reflection_v1.ServerReflectionRequest) (*grpc_reflection_v1.ServerReflectionResponse, error) {
	c.streamLock.Lock()
	defer c.streamLock.Unlock()

	err := c.stream.Send(request)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := c.stream.Recv()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response, nil
}

// 取得したfdProtoをRegisterFileする
// 依存関係の末端からRegisterする
// 依存関係が、FileA -> FileB の時、registerFuncが registerFunc(FileB), registerFunc(FileA) の順で呼ばれる
// 深さ優先探索でRegister
func (c *clientImpl) resolveFiles(
	registerFunc func(fdProto *descriptorpb.FileDescriptorProto) error,
	fdProto *descriptorpb.FileDescriptorProto,
) error {
	return c.resolveFilesInner(registerFunc, fdProto)
}

func (c *clientImpl) resolveFilesInner(
	registerFunc func(fdProto *descriptorpb.FileDescriptorProto) error,
	fdProto *descriptorpb.FileDescriptorProto,
) error {
	// 依存関係にあるファイルの、依存関係を解決
	for _, depFdName := range fdProto.Dependency {
		depReq := grpc_reflection_v1.ServerReflectionRequest{
			MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_FileByFilename{
				FileByFilename: depFdName,
			},
		}
		depResp, err := c.call(&depReq)
		if err != nil {
			return errors.WithStack(err)
		}

		for _, depFdProtoBytes := range depResp.GetFileDescriptorResponse().FileDescriptorProto {
			zap.L().Debug("Protofile Dependency", zap.String("sourceProto", fdProto.GetName()), zap.String("dependencyProto", depFdName))
			var depFdProto descriptorpb.FileDescriptorProto
			if err := proto.Unmarshal(depFdProtoBytes, &depFdProto); err != nil {
				return errors.WithStack(err)
			}
			if err := c.resolveFilesInner(registerFunc, &depFdProto); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	// 自分自身をRegister
	if err := registerFunc(fdProto); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

package grpcreflectiface

import (
	"sync"

	"github.com/pkg/errors"
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

type Stream = grpc.BidiStreamingClient[grpc_reflection_v1.ServerReflectionRequest, grpc_reflection_v1.ServerReflectionResponse]

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

func (c *clientImpl) resolveFiles(files map[string]*descriptorpb.FileDescriptorProto, dep string) error {
	if _, ok := files[dep]; ok {
		return nil
	}

	depReq := grpc_reflection_v1.ServerReflectionRequest{
		MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_FileByFilename{
			FileByFilename: dep,
		},
	}

	depResp, err := c.call(&depReq)
	if err != nil {
		// Skip if file not found
		return nil
	}

	for _, depFdProtoBytes := range depResp.GetFileDescriptorResponse().FileDescriptorProto {
		var depFdProto descriptorpb.FileDescriptorProto
		if err := proto.Unmarshal(depFdProtoBytes, &depFdProto); err != nil {
			return errors.WithStack(err)
		}
		files[depFdProto.GetName()] = &depFdProto
		// Resolve nested dependencies
		for _, nestedDep := range depFdProto.GetDependency() {
			if err := c.resolveFiles(files, nestedDep); err != nil {
				return err
			}
		}
	}

	return nil
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

	files := map[string]*descriptorpb.FileDescriptorProto{}
	for _, fdProtoBytes := range resp.GetFileDescriptorResponse().FileDescriptorProto {
		var fdProto descriptorpb.FileDescriptorProto
		if err := proto.Unmarshal(fdProtoBytes, &fdProto); err != nil {
			return nil, errors.WithStack(err)
		}
		files[fdProto.GetName()] = &fdProto

		// 依存するprotoファイルを取得
		for _, dep := range fdProto.GetDependency() {
			if err := c.resolveFiles(files, dep); err != nil {
				return nil, err
			}
		}
	}

	// ファイルをレジストリに登録
	registry := &protoregistry.Files{}
	for _, fdProto := range files {
		file, err := protodesc.NewFile(fdProto, registry)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if err := registry.RegisterFile(file); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// サービスディスクリプタを探す
	for _, fdProto := range files {
		file, err := protodesc.NewFile(fdProto, registry)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		svcDescriptor = file.Services().ByName(protoreflect.FullName(serviceName).Name())
		if svcDescriptor != nil {
			break
		}
	}

	return svcDescriptor, nil
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

	var dynamicMessage proto.Message
	files := map[string]*descriptorpb.FileDescriptorProto{}
	for _, fdProtoBytes := range resp.GetFileDescriptorResponse().FileDescriptorProto {
		var fdProto descriptorpb.FileDescriptorProto

		if err := proto.Unmarshal(fdProtoBytes, &fdProto); err != nil {
			return nil, errors.WithStack(err)
		}
		files[fdProto.GetName()] = &fdProto

		// 依存するprotoファイルを取得
		for _, dep := range fdProto.GetDependency() {
			if err := c.resolveFiles(files, dep); err != nil {
				return nil, err
			}
		}
	}

	// ファイルをレジストリに登録
	registry := &protoregistry.Files{}
	for _, fdProto := range files {
		file, err := protodesc.NewFile(fdProto, registry)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if err := registry.RegisterFile(file); err != nil {
			return nil, errors.WithStack(err)
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

package grpcreflectiface

import (
	"sync"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
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
	ResolveService(serviceName string) (*desc.ServiceDescriptor, error)
	ResolveMessage(messageName string) (proto.Message, error)
}

type Stream = grpc.BidiStreamingClient[grpc_reflection_v1.ServerReflectionRequest, grpc_reflection_v1.ServerReflectionResponse]

type clientImpl struct {
	rawGrpcReflectClient *grpcreflect.Client
	stream               Stream

	streamLock sync.Mutex
}

func NewClient(stream Stream, rawGrpcReflectClient *grpcreflect.Client) Client {
	return &clientImpl{
		rawGrpcReflectClient: rawGrpcReflectClient,
		stream:               stream,
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

func (c *clientImpl) ResolveService(serviceName string) (*desc.ServiceDescriptor, error) {
	return c.rawGrpcReflectClient.ResolveService(serviceName)
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
	for _, fdProtoBytes := range resp.GetFileDescriptorResponse().FileDescriptorProto {
		var fdProto descriptorpb.FileDescriptorProto

		if err := proto.Unmarshal(fdProtoBytes, &fdProto); err != nil {
			return nil, errors.WithStack(err)
		}

		file, err := protodesc.NewFile(&fdProto, protoregistry.GlobalFiles)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		msgDescriptor := file.Messages().ByName(protoreflect.FullName(messageName).Name())
		if msgDescriptor == nil {
			continue
		}

		// dynamicpbを使用してメッセージを動的に生成
		dynamicMessage = dynamicpb.NewMessage(msgDescriptor)
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

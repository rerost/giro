package grpcreflectiface

import (
	"sync"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1"
)

type Client interface {
	ListServices() ([]string, error)
	ResolveService(serviceName string) (*desc.ServiceDescriptor, error)
	ResolveMessage(messageName string) (*desc.MessageDescriptor, error)
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

func (c *clientImpl) ResolveMessage(messageName string) (*desc.MessageDescriptor, error) {
	return c.rawGrpcReflectClient.ResolveMessage(messageName)
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

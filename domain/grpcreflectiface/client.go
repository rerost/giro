//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source=./client.go -destination=../../mock/grpcreflectiface/client_test.go -package=mockgrpcreflectiface

package grpcreflectiface

import (
	"context"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/pkg/errors"
	"google.golang.org/grpc/reflection/grpc_reflection_v1"
)

type Client interface {
	ListServices() ([]string, error)
	ResolveService(serviceName string) (*desc.ServiceDescriptor, error)
	ResolveMessage(messageName string) (*desc.MessageDescriptor, error)
}

type clientImpl struct {
	rawClient            grpc_reflection_v1.ServerReflectionClient
	rawGrpcReflectClient *grpcreflect.Client
}

func NewClient(rawClient grpc_reflection_v1.ServerReflectionClient, rawGrpcReflectClient *grpcreflect.Client) Client {
	return &clientImpl{
		rawClient:            rawClient,
		rawGrpcReflectClient: rawGrpcReflectClient,
	}
}

func (c *clientImpl) ListServices() ([]string, error) {
	stream, err := c.rawClient.ServerReflectionInfo(context.TODO())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stream.CloseSend()

	err = stream.Send(&grpc_reflection_v1.ServerReflectionRequest{
		MessageRequest: &grpc_reflection_v1.ServerReflectionRequest_ListServices{
			ListServices: "*",
		},
	})

	if err != nil {
		return nil, errors.WithStack(err)
	}

	services := []string{}
	response, err := stream.Recv()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	recvServices := response.GetListServicesResponse().Service
	for _, service := range recvServices {
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

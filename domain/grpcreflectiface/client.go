//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source=./client.go -destination=../../mock/grpcreflectiface/client_test.go -package=mockgrpcreflectiface

package grpcreflectiface

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
)

type Client interface {
	ListServices() ([]string, error)
	ResolveService(serviceName string) (*desc.ServiceDescriptor, error)
	ResolveMessage(messageName string) (*desc.MessageDescriptor, error)
}

type clientImpl struct {
	rawClient *grpcreflect.Client
}

func NewClient(rawClient *grpcreflect.Client) Client {
	return &clientImpl{
		rawClient: rawClient,
	}
}

func (c *clientImpl) ListServices() ([]string, error) {
	return c.rawClient.ListServices()
}

func (c *clientImpl) ResolveService(serviceName string) (*desc.ServiceDescriptor, error) {
	return c.rawClient.ResolveService(serviceName)
}

func (c *clientImpl) ResolveMessage(messageName string) (*desc.MessageDescriptor, error) {
	return c.rawClient.ResolveMessage(messageName)
}

package grpcreflectiface

import "github.com/jhump/protoreflect/desc"

type Client interface {
	ListServices() ([]string, error)
	ResolveService(serviceName string) (*desc.ServiceDescriptor, error)
	ResolveMessage(messageName string) (*desc.MessageDescriptor, error)
}

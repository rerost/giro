package messagename

import (
	"context"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/pkg/errors"
)

type MessageName string

type MessageNameResolver interface {
	RequestMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error)
	ResponseMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error)
}

type messageNameResolverImpl struct {
	client *grpcreflect.Client
}

func NewMessageNameResolver(client *grpcreflect.Client) MessageNameResolver {
	return &messageNameResolverImpl{
		client: client,
	}
}

func (mnr *messageNameResolverImpl) RequestMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error) {
	md, err := mnr.resolveMethodDescriptor(ctx, serviceName, methodName)
	if err != nil {
		return "", errors.WithStack(err)
	}

	messageDescriptor := md.GetInputType()
	return MessageName(messageDescriptor.GetFullyQualifiedName()), nil
}

func (mnr *messageNameResolverImpl) ResponseMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error) {
	md, err := mnr.resolveMethodDescriptor(ctx, serviceName, methodName)
	if err != nil {
		return "", errors.WithStack(err)
	}

	messageDescriptor := md.GetOutputType()
	return MessageName(messageDescriptor.GetFullyQualifiedName()), nil
}

func (mnr *messageNameResolverImpl) resolveMethodDescriptor(ctx context.Context, serviceName string, methodName string) (*desc.MethodDescriptor, error) {
	sd, err := mnr.client.ResolveService(serviceName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	md := sd.FindMethodByName(methodName)
	return md, nil
}

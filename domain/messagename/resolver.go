package messagename

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rerost/giro/domain/grpcreflectiface"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type MessageName string

type MessageNameResolver interface {
	RequestMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error)
	ResponseMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error)
}

type messageNameResolverImpl struct {
	client grpcreflectiface.Client
}

func NewMessageNameResolver(client grpcreflectiface.Client) MessageNameResolver {
	return &messageNameResolverImpl{
		client: client,
	}
}

var (
	MethodNotFoundError = errors.New("Method not Found")
)

func (mnr *messageNameResolverImpl) RequestMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error) {
	md, err := mnr.resolveMethodDescriptor(ctx, serviceName, methodName)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return MessageName(md.Input().FullName()), nil
}

func (mnr *messageNameResolverImpl) ResponseMessageName(ctx context.Context, serviceName string, methodName string) (MessageName, error) {
	md, err := mnr.resolveMethodDescriptor(ctx, serviceName, methodName)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return MessageName(md.Output().FullName()), nil
}

func (mnr *messageNameResolverImpl) resolveMethodDescriptor(ctx context.Context, serviceName string, methodName string) (protoreflect.MethodDescriptor, error) {
	sd, err := mnr.client.ResolveService(serviceName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	md := sd.Methods().ByName(protoreflect.FullName(methodName).Name())
	if md == nil {
		return nil, errors.WithStack(MethodNotFoundError)
	}
	return md, nil
}

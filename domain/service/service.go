package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rerost/giro/domain/grpcreflectiface"
	"github.com/rerost/giro/domain/host"
	"github.com/rerost/giro/domain/message"
	"github.com/rerost/giro/domain/messagename"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ServiceService interface {
	Call(ctx context.Context, serviceName string, methodName string, metadata map[string]string, body message.JSON) (message.JSON, error)
	Ls(ctx context.Context, serviceName *string, methodName *string) ([]Service, error)
}

type Service struct {
	Name        string
	MethodNames []string
}

type serviceServiceImpl struct {
	hostResolver        host.HostResolver
	grpcreflectClient   grpcreflectiface.Client
	messageService      message.MessageService
	messageNameResolver messagename.MessageNameResolver
	grpcClientOpts      []grpc.DialOption
}

func NewServiceService(grpcreflectClient grpcreflectiface.Client, hostResolver host.HostResolver, messageNameResolver messagename.MessageNameResolver, messageService message.MessageService) ServiceService {
	return &serviceServiceImpl{
		hostResolver:        hostResolver,
		messageNameResolver: messageNameResolver,
		messageService:      messageService,
		grpcreflectClient:   grpcreflectClient,
	}
}

func (ss *serviceServiceImpl) Call(ctx context.Context, serviceName string, methodName string, md map[string]string, body message.JSON) (message.JSON, error) {
	grpcClient, err := ss.NewClient(ctx, serviceName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	requestMessageName, err := ss.messageNameResolver.RequestMessageName(ctx, serviceName, methodName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	requestDynamicMessage, err := ss.messageService.ToDynamicMessage(ctx, requestMessageName, body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	responseMessageName, err := ss.messageNameResolver.ResponseMessageName(ctx, serviceName, methodName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	responseDynamicMessage, err := ss.messageService.ToDynamicMessage(ctx, responseMessageName, message.JSON("{}"))

	cctx := ctx
	for k, v := range md {
		k := k
		v := v
		cctx = metadata.AppendToOutgoingContext(cctx, k, v)
	}
	err = grpcClient.Invoke(cctx, ss.fullMethodName(serviceName, methodName), requestDynamicMessage, responseDynamicMessage)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ss.messageService.DynamicMessageToJSON(ctx, responseDynamicMessage)
}

func (ss *serviceServiceImpl) Ls(ctx context.Context, serviceName *string, methodName *string) ([]Service, error) {
	if serviceName == nil && methodName == nil {
		svc := []Service{}
		serviceNames, err := ss.grpcreflectClient.ListServices()
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, sn := range serviceNames {
			svc = append(svc, Service{Name: sn})
		}
		return svc, nil
	} else if serviceName != nil && methodName == nil {
		svc := Service{Name: *serviceName}
		sd, err := ss.grpcreflectClient.ResolveService(*serviceName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, md := range sd.GetMethods() {
			svc.MethodNames = append(svc.MethodNames, md.GetName())
		}

		return []Service{svc}, nil
	} else if serviceName != nil && methodName != nil {
		svc := Service{Name: *serviceName}
		sd, err := ss.grpcreflectClient.ResolveService(*serviceName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for _, md := range sd.GetMethods() {
			if *methodName == md.GetName() {
				svc.MethodNames = append(svc.MethodNames, md.GetName())

				return []Service{svc}, nil
			}
		}

		return nil, errors.New("Method not found")
	}

	panic("Unexpected")
}

func (ss *serviceServiceImpl) NewClient(ctx context.Context, serviceName string) (*grpc.ClientConn, error) {
	if ss.hostResolver == nil {
		return nil, errors.New("Unsupported")
	}
	target, err := ss.hostResolver.Resolve(ctx, serviceName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	conn, err := grpc.DialContext(ctx, string(target), grpc.WithInsecure())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return conn, nil
}

func (ss *serviceServiceImpl) fullMethodName(serviceName string, methodName string) string {
	return fmt.Sprintf("/%s/%s", serviceName, methodName)
}

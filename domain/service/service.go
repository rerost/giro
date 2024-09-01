package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rerost/giro/domain/grpcreflectiface"
	"github.com/rerost/giro/domain/host"
	"github.com/rerost/giro/domain/message"
	"github.com/rerost/giro/domain/messagename"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

type ServiceService interface {
	Call(ctx context.Context, serviceName string, methodName string, metadata map[string]string, body message.JSON, callTimeout time.Duration) (message.JSON, error)
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
}

func NewServiceService(grpcreflectClient grpcreflectiface.Client, hostResolver host.HostResolver, messageNameResolver messagename.MessageNameResolver, messageService message.MessageService) ServiceService {
	return &serviceServiceImpl{
		hostResolver:        hostResolver,
		messageNameResolver: messageNameResolver,
		messageService:      messageService,
		grpcreflectClient:   grpcreflectClient,
	}
}

func (ss *serviceServiceImpl) Call(ctx context.Context, serviceName string, methodName string, md map[string]string, body message.JSON, callTimeout time.Duration) (message.JSON, error) {
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
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cctx := ctx
	for k, v := range md {
		k := k
		v := v
		cctx = metadata.AppendToOutgoingContext(cctx, k, v)
	}

	if callTimeout != 0 {
		var cancel context.CancelFunc
		cctx, cancel = context.WithTimeout(cctx, callTimeout)
		defer cancel()
	}
	err = grpcClient.Invoke(cctx, ss.fullMethodName(serviceName, methodName), requestDynamicMessage, responseDynamicMessage)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	json, err := protojson.Marshal(responseDynamicMessage)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return json, nil
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
		if sd == nil {
			return []Service{svc}, nil
		}

		for i := 0; i < sd.Methods().Len(); i++ {
			md := sd.Methods().Get(i)
			svc.MethodNames = append(svc.MethodNames, string(md.Name()))
		}

		return []Service{svc}, nil
	} else if serviceName != nil && methodName != nil {
		svc := Service{Name: *serviceName}
		sd, err := ss.grpcreflectClient.ResolveService(*serviceName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for i := 0; i < sd.Methods().Len(); i++ {
			md := sd.Methods().Get(i)
			if *methodName == string(md.Name()) {
				svc.MethodNames = append(svc.MethodNames, string(md.FullName()))
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
	conn, err := grpc.NewClient(string(target), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return conn, nil
}

func (ss *serviceServiceImpl) fullMethodName(serviceName string, methodName string) string {
	return fmt.Sprintf("/%s/%s", serviceName, methodName)
}

package host

import (
	"context"

	"github.com/pkg/errors"
	hosts_pb "github.com/rerost/giro/pb/hosts"
)

type HostResolver interface {
	Resolve(ctx context.Context, serviceName string) (string, error)
}

type hostResolverImpl struct {
	client hosts_pb.HostServiceClient
}

func NewHostResolver(client hosts_pb.HostServiceClient) HostResolver {
	return &hostResolverImpl{
		client: client,
	}
}

func (hr hostResolverImpl) Resolve(ctx context.Context, serviceName string) (string, error) {
	res, err := hr.client.GetHost(ctx, &hosts_pb.GetHostRequest{
		ServiceName: serviceName,
	})

	if err != nil {
		return "", errors.WithStack(err)
	}

	return res.GetHost(), nil
}

type constHostResolverImpl struct {
	addr string
}

func NewConstHostResolver(addr string) HostResolver {
	return &constHostResolverImpl{
		addr: addr,
	}
}

func (chr *constHostResolverImpl) Resolve(ctx context.Context, serviceName string) (string, error) {
	return chr.addr, nil
}

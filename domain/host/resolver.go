package host

import "context"

type HostResolver interface {
	Resolve(ctx context.Context, serviceName string) (string, error)
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

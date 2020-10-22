package dummyserver

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	hosts_pb "github.com/rerost/giro/pb/hosts"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
)

func Run(port string) (func(), error) {
	return runServer(port)
}

type testServiceServerImpl struct{}

func newTestService() TestServiceServer {
	return &testServiceServerImpl{}
}

func (s *testServiceServerImpl) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	md := map[string]*MetadataValue{}

	metadata, ok := metadata.FromIncomingContext(ctx)
	if ok {
		_metadata := map[string][]string(metadata)

		for k, v := range _metadata {
			md[k] = &MetadataValue{
				Value: v,
			}
		}
	}

	return &EchoResponse{
		Message: req.GetMessage(),
		Metadata: &Metadata{
			Metadata: md,
		},
	}, nil
}

func NewHostsServiceServer() hosts_pb.HostServiceServer {
	return &hostsServiceServerImpl{
		hosts: map[string]string{
			"rerost.giro.v1.TestService": "localhost:5000",
		},
	}
}

type hostsServiceServerImpl struct {
	hosts map[string]string
}

func (s *hostsServiceServerImpl) ListHosts(_ context.Context, req *hosts_pb.ListHostsRequest) (*hosts_pb.ListHostsResponse, error) {
	serviceName := req.GetServiceName()
	host, ok := s.hosts[serviceName]
	if !ok {
		return nil, status.Error(codes.NotFound, "NotFound")
	}

	return &hosts_pb.ListHostsResponse{
		Host: host,
	}, nil
}

func runServer(port string) (func(), error) {
	log.Printf("listen: %v\n", port)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return nil, errors.WithStack(err)
	}

	server := grpc.NewServer()
	RegisterTestServiceServer(server, newTestService())
	hosts_pb.RegisterHostServiceServer(server, NewHostsServiceServer())
	reflection.Register(server)
	go func() {
		server.Serve(lis)
	}()
	return func() {
		lis.Close()
	}, nil
}

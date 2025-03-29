package dummyserver

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	hosts_pb "github.com/rerost/giro/rerost/giro"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func Run(port string) (func(), error) {
	return runServer(port)
}

type testServiceServerImpl struct {
	UnimplementedTestServiceServer
}

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

func (s *testServiceServerImpl) EmptyCall(ctx context.Context, _ *emptypb.Empty) (*EmptyResponse, error) {
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

	return &EmptyResponse{
		Status: "ok",
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
	hosts_pb.UnimplementedHostServiceServer

	hosts map[string]string
}

func (s *hostsServiceServerImpl) GetHost(_ context.Context, req *hosts_pb.GetHostRequest) (*hosts_pb.GetHostResponse, error) {
	serviceName := req.GetServiceName()
	host, ok := s.hosts[serviceName]
	if !ok {
		return nil, status.Error(codes.NotFound, "NotFound")
	}

	return &hosts_pb.GetHostResponse{
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

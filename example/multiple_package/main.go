package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	github_com_rerost_giro_example_multiple_package_protos_one "github.com/rerost/giro/example/multiple_package/protos/one"
	github_com_rerost_giro_example_multiple_package_protos_two "github.com/rerost/giro/example/multiple_package/protos/two"
	hosts_pb "github.com/rerost/giro/rerost/giro"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func NewHostsServiceServer() hosts_pb.HostServiceServer {
	return &hostsServiceServerImpl{
		hosts: map[string]string{},
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
func Newgithub_com_rerost_giro_example_multiple_package_protos_twoBqvService() github_com_rerost_giro_example_multiple_package_protos_two.BqvServiceServer {
	return &github_com_rerost_giro_example_multiple_package_protos_twoBqvServiceImpl{}
}

type github_com_rerost_giro_example_multiple_package_protos_twoBqvServiceImpl struct {
	github_com_rerost_giro_example_multiple_package_protos_two.UnimplementedBqvServiceServer
}

func Newgithub_com_rerost_giro_example_multiple_package_protos_oneGiroService() github_com_rerost_giro_example_multiple_package_protos_one.GiroServiceServer {
	return &github_com_rerost_giro_example_multiple_package_protos_oneGiroServiceImpl{}
}

type github_com_rerost_giro_example_multiple_package_protos_oneGiroServiceImpl struct {
	github_com_rerost_giro_example_multiple_package_protos_one.UnimplementedGiroServiceServer
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		fmt.Println("Please set APP_PORT")
		port = "5000"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	server := grpc.NewServer()
	healthpb.RegisterHealthServer(server, health.NewServer())
	github_com_rerost_giro_example_multiple_package_protos_two.RegisterBqvServiceServer(server, Newgithub_com_rerost_giro_example_multiple_package_protos_twoBqvService())
	github_com_rerost_giro_example_multiple_package_protos_one.RegisterGiroServiceServer(server, Newgithub_com_rerost_giro_example_multiple_package_protos_oneGiroService())
	hosts_pb.RegisterHostServiceServer(server, NewHostsServiceServer())
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

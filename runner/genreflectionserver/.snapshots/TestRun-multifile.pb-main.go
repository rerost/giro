package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	hosts_pb "github.com/rerost/giro/rerost/giro"
	github_com_rerost_giro_runner_genreflectionserver_testdata_onefile "github.com/rerost/giro/runner/genreflectionserver/testdata/onefile"
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
	github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.RegisterGiroServiceServer(server, github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.UnimplementedGiroServiceServer{})
	github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.RegisterBqvServiceServer(server, github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.UnimplementedBqvServiceServer{})
	hosts_pb.RegisterHostServiceServer(server, NewHostsServiceServer())
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


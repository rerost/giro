package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	github_com_rerost_giro_pb_hosts "github.com/rerost/giro/pb/hosts"
	hosts_pb "github.com/rerost/giro/pb/hosts"
	github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_option "github.com/rerost/giro/runner/genreflectionserver/testdata/with_host_option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func NewHostsServiceServer() hosts_pb.HostServiceServer {
	return &hostsServiceServerImpl{
		hosts: map[string]string{
			"runner.genreflectionserver.testdata.onefile.GiroService": "localhost:5000",
		},
	}
}

type hostsServiceServerImpl struct {
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
func Newgithub_com_rerost_giro_pb_hostsHostService() github_com_rerost_giro_pb_hosts.HostServiceServer {
	return &github_com_rerost_giro_pb_hostsHostServiceImpl{}
}

type github_com_rerost_giro_pb_hostsHostServiceImpl struct {
}

func (s *github_com_rerost_giro_pb_hostsHostServiceImpl) ListHosts(ctx context.Context, req *github_com_rerost_giro_pb_hosts.ListHostsRequest) (*github_com_rerost_giro_pb_hosts.ListHostsResponse, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "Dummy")
}
func Newgithub_com_rerost_giro_runner_genreflectionserver_testdata_with_host_optionGiroService() github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_option.GiroServiceServer {
	return &github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_optionGiroServiceImpl{}
}

type github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_optionGiroServiceImpl struct {
}

func (s *github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_optionGiroServiceImpl) GiroTest1(ctx context.Context, req *github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_option.GiroTestRequest1) (*github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_option.GiroTestResponse1, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "Dummy")
}
func (s *github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_optionGiroServiceImpl) GiroTest2(ctx context.Context, req *github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_option.GiroTestRequest2) (*github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_option.GiroTestResponse2, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "Dummy")
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
	github_com_rerost_giro_runner_genreflectionserver_testdata_with_host_option.RegisterGiroServiceServer(server, Newgithub_com_rerost_giro_runner_genreflectionserver_testdata_with_host_optionGiroService())
	hosts_pb.RegisterHostServiceServer(server, NewHostsServiceServer())
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


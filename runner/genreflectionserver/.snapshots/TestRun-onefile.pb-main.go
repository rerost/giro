package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	hosts_pb "github.com/rerost/giro/pb/hosts"
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
func Newgithub_com_rerost_giro_runner_genreflectionserver_testdata_onefileGiroService() github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.GiroServiceServer {
	return &github_com_rerost_giro_runner_genreflectionserver_testdata_onefileGiroServiceImpl{}
}

type github_com_rerost_giro_runner_genreflectionserver_testdata_onefileGiroServiceImpl struct {
}

func (s *github_com_rerost_giro_runner_genreflectionserver_testdata_onefileGiroServiceImpl) GiroTest1(ctx context.Context, req *github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.GiroTestRequest1) (*github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.GiroTestResponse1, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "Dummy")
}
func (s *github_com_rerost_giro_runner_genreflectionserver_testdata_onefileGiroServiceImpl) GiroTest2(ctx context.Context, req *github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.GiroTestRequest2) (*github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.GiroTestResponse2, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "Dummy")
}
func Newgithub_com_rerost_giro_runner_genreflectionserver_testdata_onefileBqvService() github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.BqvServiceServer {
	return &github_com_rerost_giro_runner_genreflectionserver_testdata_onefileBqvServiceImpl{}
}

type github_com_rerost_giro_runner_genreflectionserver_testdata_onefileBqvServiceImpl struct {
}

func (s *github_com_rerost_giro_runner_genreflectionserver_testdata_onefileBqvServiceImpl) BqvTest1(ctx context.Context, req *github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.BqvTestRequest1) (*github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.BqvTestResponse1, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "Dummy")
}
func (s *github_com_rerost_giro_runner_genreflectionserver_testdata_onefileBqvServiceImpl) BqvTest2(ctx context.Context, req *github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.BqvTestRequest2) (*github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.BqvTestResponse2, error) {
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
	github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.RegisterGiroServiceServer(server, Newgithub_com_rerost_giro_runner_genreflectionserver_testdata_onefileGiroService())
	github_com_rerost_giro_runner_genreflectionserver_testdata_onefile.RegisterBqvServiceServer(server, Newgithub_com_rerost_giro_runner_genreflectionserver_testdata_onefileBqvService())
	hosts_pb.RegisterHostServiceServer(server, NewHostsServiceServer())
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


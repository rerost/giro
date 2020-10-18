package dummyserver

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run(port string) (func(), error) {
	return runServer(port)
}

type testServiceServerImpl struct{}

func newTestService() TestServiceServer {
	return &testServiceServerImpl{}
}

func (s *testServiceServerImpl) Echo(_ context.Context, req *EchoRequest) (*EchoResponse, error) {
	return &EchoResponse{
		Message: req.GetMessage(),
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
	reflection.Register(server)
	go func() {
		server.Serve(lis)
	}()
	return func() {
		lis.Close()
	}, nil
}

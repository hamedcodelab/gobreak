package main

import (
	"context"
	"fmt"
	helloworld "github.com/hamedcodelab/gobreak/example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// Server struct
type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements the GreeterServer interface.
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	// Create a listener on TCP port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the Greeter service
	helloworld.RegisterGreeterServer(s, &server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Start serving
	fmt.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

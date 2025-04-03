package main

import (
	"context"
	"fmt"
	helloworld "github.com/hamedcodelab/gobreak/example/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a Greeter client
	c := helloworld.NewGreeterClient(conn)

	// Make a request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the SayHello RPC
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print the response
	fmt.Printf("Greeting: %s\n", r.GetMessage())
}

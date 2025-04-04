package main

import (
	"context"
	"fmt"
	"github.com/hamedcodelab/gobreak"
	helloworld "github.com/hamedcodelab/gobreak/example/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// only create gobreak
	gbrk := gobreak.NewBreaker(gobreak.WithFailureThreshold(1), gobreak.WithRecoveryTime(time.Minute*4), gobreak.WithHalfOpenMaxRequests(2))

	var conn *grpc.ClientConn
	// first grpc call with error
	gbrk.Execute(func() error {
		// Connect to the server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var err error
		conn, err = grpc.DialContext(ctx, ":50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		return nil
	})

	// first grpc call again error
	gbrk.Execute(func() error {
		// Connect to the server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var err error
		conn, err = grpc.DialContext(ctx, ":50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		return nil
	})

	time.Sleep(time.Second * 5)

	// first grpc call again error
	gbrk.Execute(func() error {
		// Connect to the server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		var err error
		conn, err = grpc.DialContext(ctx, ":50051", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		return nil
	})

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

package main

import (
	"context"
	"log"

	"github.com/youshy/gRPC-API/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("Client is up!\n")

	// WithInsecure overrides default ssl
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not establish connection %v\n", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Artur",
			LastName:  "Kondas",
		},
	}

	ctx := context.Background()

	res, err := c.Greet(ctx, req)
	if err != nil {
		log.Fatalf("error calling Greet RPC %v\n", err)
	}
	log.Printf("Response from greet: %v\n", res.Result)
}

package main

import (
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
	log.Printf("Created client %f\n", c)
}

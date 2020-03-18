package main

import (
	"log"
	"net"

	"github.com/youshy/gRPC-API/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.Printf("Hello!\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}

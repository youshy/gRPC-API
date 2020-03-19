package main

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/youshy/gRPC-API/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greeter invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName + "!"
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("Greet stream invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		// I know, it's not the nicest, but it's something
		time.Sleep(time.Second)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet invoked\n")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// stream finished
			res := &greetpb.LongGreetResponse{
				Result: result,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			log.Fatalf("error while reading stream %v\n", err)
		}

		firstName := req.GetGreeting().GetFirstName()
		result += firstName + " "

	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone invoked\n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// stream finished
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream %v\n", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		res := &greetpb.GreetEveryoneResponse{
			Result: result,
		}
		err = stream.Send(res)
		if err != nil {
			log.Fatalf("error while streaming data to client %v\n", err)
		}
	}
}

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

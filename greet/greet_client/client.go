package main

import (
	"context"
	"io"
	"log"
	"time"

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

	// doUnary(c)
	//	doServerStreaming(c)
	doClientStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	log.Printf("Started server streaming!\n")

	ctx := context.Background()

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Artur",
			LastName:  "Kondas",
		},
	}

	resStream, err := c.GreetManyTimes(ctx, req)
	if err != nil {
		log.Fatalf("Error while streaming: %v\n", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream %v\n", err)
		}
		log.Printf("Greet many times: %v\n", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	log.Printf("Started client streaming\n")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tom",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jerry",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tweetie",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Garfield",
			},
		},
	}

	ctx := context.Background()

	stream, err := c.LongGreet(ctx)
	if err != nil {
		log.Fatalf("error while calling LongGreet %v\n", err)
	}

	for iter, req := range requests {
		// this is simulating bigger stuff
		log.Printf("Sending chunk no %v\n", iter)
		stream.Send(req)
		time.Sleep(time.Second)

	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet %v\n", err)
	}

	log.Printf("LongGreet Response: %v\n", res)
}

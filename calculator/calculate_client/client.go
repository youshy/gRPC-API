package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/youshy/gRPC-API/calculator/calculatepb"
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

	c := calculatepb.NewCalculateServiceClient(conn)

	/*
		Sum(c, 10, 3)
		Division(c, 120, 423)
		Multiply(c, 25, 246524)
		Substract(c, 234542, 534)
	*/

	// PrimeNumberStream(c, 12390392840)

	// StreamCalcuateAverage(c)

	// BiDiFindMaximumStream(c)

}

// Unary
func Sum(c calculatepb.CalculateServiceClient, first, second int64) {
	req := &calculatepb.CalculateRequest{
		Calculate: &calculatepb.Calculate{
			FirstNumber:  first,
			SecondNumber: second,
		},
	}

	ctx := context.Background()

	res, err := c.CalculateSum(ctx, req)
	if err != nil {
		log.Fatalf("error calling Calc RPC %v\n", err)
	}
	log.Printf("Sum: %v\n", res.Result)
}

func Division(c calculatepb.CalculateServiceClient, first, second int64) {
	req := &calculatepb.CalculateRequest{
		Calculate: &calculatepb.Calculate{
			FirstNumber:  first,
			SecondNumber: second,
		},
	}

	ctx := context.Background()

	res, err := c.CalculateDivision(ctx, req)
	if err != nil {
		log.Fatalf("error calling Calc RPC %v\n", err)
	}
	log.Printf("Division: %v\n", res.Result)
}

func Multiply(c calculatepb.CalculateServiceClient, first, second int64) {
	req := &calculatepb.CalculateRequest{
		Calculate: &calculatepb.Calculate{
			FirstNumber:  first,
			SecondNumber: second,
		},
	}

	ctx := context.Background()

	res, err := c.CalculateMultiply(ctx, req)
	if err != nil {
		log.Fatalf("error calling Calc RPC %v\n", err)
	}
	log.Printf("Multiply: %v\n", res.Result)
}

func Substract(c calculatepb.CalculateServiceClient, first, second int64) {
	req := &calculatepb.CalculateRequest{
		Calculate: &calculatepb.Calculate{
			FirstNumber:  first,
			SecondNumber: second,
		},
	}

	ctx := context.Background()

	res, err := c.CalculateSubstract(ctx, req)
	if err != nil {
		log.Fatalf("error calling Calc RPC %v\n", err)
	}
	log.Printf("Substract: %v\n", res.Result)
}

// Server stream
func PrimeNumberStream(c calculatepb.CalculateServiceClient, number int64) {
	ctx := context.Background()

	req := &calculatepb.PrimeNumberRequest{
		Primenumber: &calculatepb.PrimeNumber{
			Number: number,
		},
	}

	resStream, err := c.PrimeNumberDecompose(ctx, req)
	if err != nil {
		log.Fatalf("error while streaming %v\n", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("error reading stream %v\n", err)
		}
		log.Printf("Res: %v\n", msg.GetResult())
	}
}

// Client stream
func StreamCalcuateAverage(c calculatepb.CalculateServiceClient) {
	ctx := context.Background()

	requests := []*calculatepb.CalculateAverageRequest{
		&calculatepb.CalculateAverageRequest{
			Number: 3,
		},
		&calculatepb.CalculateAverageRequest{
			Number: 6,
		},
		&calculatepb.CalculateAverageRequest{
			Number: 32,
		},
		&calculatepb.CalculateAverageRequest{
			Number: 2,
		},
		&calculatepb.CalculateAverageRequest{
			Number: 69,
		},
	}

	stream, err := c.CalculateAverage(ctx)
	if err != nil {
		log.Fatalf("error while calling CalculateAverage %v\n", err)
	}

	for iter, req := range requests {
		log.Printf("Sending chunk no %v\tCalculating %v\n", iter, req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from CalculateAverage %v\n", err)
	}

	log.Printf("CalculateAverage response: %v\n", res)
}

// BiDi Stream
func BiDiFindMaximumStream(c calculatepb.CalculateServiceClient) {
	ctx := context.Background()

	stream, err := c.FindMaximum(ctx)

	if err != nil {
		log.Fatalf("error while opening stream and calling FindMaximum %v\n", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		numbers := []int64{2, 4, 6, 8, 1, 5, 94, 1, 435, 444, 432, 605}
		for _, number := range numbers {
			req := &calculatepb.FindMaximumRequest{
				Number: number,
			}
			stream.Send(req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()
	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("problem while reading the server stream %v\n", err)
				break
			}
			maximum := res.GetMaximum()
			log.Printf("Received a new maximum of %v\n", maximum)
		}
		close(waitc)
	}()

	<-waitc

}

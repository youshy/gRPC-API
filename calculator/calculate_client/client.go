package main

import (
	"context"
	"io"
	"log"

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
	PrimeNumberStream(c, 1200)
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

// Stream
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

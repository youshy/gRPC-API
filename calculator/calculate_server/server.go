package main

import (
	"context"
	"log"
	"net"

	"github.com/youshy/gRPC-API/calculator/calculatepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) CalculateSum(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	log.Printf("calculate sum invoked with %v\n", req)
	f1 := req.GetCalculate().GetFirstNumber()
	f2 := req.GetCalculate().GetSecondNumber()
	result := f1 + f2
	res := &calculatepb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func (*server) CalculateSubstract(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	log.Printf("calculate substract invoked with %v\n", req)
	f1 := req.GetCalculate().GetFirstNumber()
	f2 := req.GetCalculate().GetSecondNumber()
	result := f1 - f2
	res := &calculatepb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func (*server) CalculateMultiply(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	log.Printf("calculate multiply invoked with %v\n", req)
	f1 := req.GetCalculate().GetFirstNumber()
	f2 := req.GetCalculate().GetSecondNumber()
	result := f1 * f2
	res := &calculatepb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func (*server) CalculateDivision(ctx context.Context, req *calculatepb.CalculateRequest) (*calculatepb.CalculateResponse, error) {
	log.Printf("calculate division invoked with %v\n", req)
	f1 := req.GetCalculate().GetFirstNumber()
	f2 := req.GetCalculate().GetSecondNumber()
	result := f1 / f2
	res := &calculatepb.CalculateResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	log.Printf("Hello!\nThis is calculate grpc\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v\n", err)
	}

	s := grpc.NewServer()
	calculatepb.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}

package main

import (
	"context"
	"io"
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

func (*server) PrimeNumberDecompose(req *calculatepb.PrimeNumberRequest, stream calculatepb.CalculateService_PrimeNumberDecomposeServer) error {
	log.Printf("Prime number decompose with %v\n", req)
	number := req.GetPrimenumber().GetNumber()
	k := int64(2)
	for number > 1 {
		if number%k == 0 {
			res := &calculatepb.PrimeNumberResponse{
				Result: k,
			}
			stream.Send(res)
			number = number / k
		} else {
			k = k + 1
		}
	}

	return nil
}

func (*server) CalculateAverage(stream calculatepb.CalculateService_CalculateAverageServer) error {
	log.Printf("Calculate average invoked\n")
	var res float32
	var ele float32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// stream finished
			calculated := res / ele
			result := &calculatepb.CalculateAverageResponse{
				Result: calculated,
			}
			return stream.SendAndClose(result)
		}
		if err != nil {
			log.Fatalf("error while reading stream %v\n", err)
		}

		number := req.GetNumber()
		res += number
		ele += 1
	}

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

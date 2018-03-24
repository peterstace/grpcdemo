package main

import (
	"context"
	"log"
	"net"

	"github.com/peterstace/grpcdemo/calc"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	s := grpc.NewServer()
	calc.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("Stopped serving: %v", err)
	}
}

type server struct{}

func (server) Add(ctx context.Context, nums *calc.Numbers) (*calc.Number, error) {
	n1, err := decimal.NewFromString(nums.GetNum1())
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "parsing num1: %v", err)
	}
	n2, err := decimal.NewFromString(nums.GetNum2())
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "parsing num2: %v", err)
	}
	return &calc.Number{
		Num: n1.Add(n2).String(),
	}, nil
}

package main

import (
	"context"
	"fmt"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Printf("Sum was involked with %v\n", in)

	return &pb.SumResponse{
		Result: in.Number1 + in.Number2,
	}, nil
}

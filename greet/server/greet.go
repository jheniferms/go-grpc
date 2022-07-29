package main

import (
	"context"
	"fmt"

	pb "github.com/jheniferms/go-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

package main

import (
	"context"
	"log"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient) {
	log.Println("doSqrt was involked")

	request := &pb.SqrtRequest{
		Number: -4,
	}

	response, err := c.Sqrt(context.Background(), request)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message: %s\n", e.Message())
			log.Printf("Error code: %s\n", e.Code())
		} else {
			log.Fatalf("A non gRPC error: %v\n", e)
		}
		return
	}

	log.Printf("Sqrt: %f\n", response.Result)
}

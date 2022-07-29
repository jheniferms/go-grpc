package main

import (
	"context"
	"log"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {

	log.Println("doSum was involked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Number1: 7,
		Number2: 3,
	})

	if err != nil {
		log.Fatalf("Could not get sum: %v\n", err)
	}

	log.Printf("Result %s\n", res)
}

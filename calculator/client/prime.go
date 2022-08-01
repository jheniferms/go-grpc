package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	fmt.Println("doPrime was involked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while get primes: %v", err)
	}

	for {

		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		fmt.Printf("Primes:%v\n", msg)
	}
}

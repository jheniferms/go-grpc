package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg was involked")

	reqs := []*pb.AvgRequest{
		{Number: 2},
		{Number: 3},
		{Number: 10},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending %v", req)

		stream.Send(req)

		time.Sleep(1 * time.Second)

	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}

	log.Printf("Avg: %v\n", res.Result)

}

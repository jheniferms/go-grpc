package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jheniferms/go-grpc/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was involked")

	reqs := []*pb.GreetRequest{
		{FirstName: "jheni"},
		{FirstName: "arthur"},
		{FirstName: "sophia"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)

		stream.Send(req)

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}

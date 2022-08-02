package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax was involked")

	request := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	waitClient := make(chan struct{})

	// send request
	go func() {
		for _, req := range request {
			log.Printf("Sending request %v", req)

			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// listen to server
	go func() {
		for {
			response, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving data from server: %v", err)
				return
			}

			log.Printf("Received: %v", response.Result)
		}

		close(waitClient)
	}()

	<-waitClient
}

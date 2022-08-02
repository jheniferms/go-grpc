package main

import (
	"context"
	"log"
	"time"

	pb "github.com/jheniferms/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was involked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	request := &pb.GreetRequest{
		FirstName: "jheni",
	}

	response, err := c.GreetWithDeadline(ctx, request)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error:%v", err)
		}
		return
	}

	log.Printf("GreetWithDeadline: %s\n", response.Result)
}

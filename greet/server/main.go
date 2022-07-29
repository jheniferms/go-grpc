package main

import (
	"log"
	"net"

	pb "github.com/jheniferms/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

var adrr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", adrr)

	if err != nil {
		log.Fatalf("Failed to listen to  on: %v\n", err)
	}

	log.Printf("Listening on %s\n", adrr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}

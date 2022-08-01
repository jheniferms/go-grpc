package main

import (
	"io"
	"log"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg was involked")

	var countNumber float32 = 0
	var sumNumber float32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: sumNumber / countNumber,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		sumNumber += req.Number
		countNumber++
	}
}

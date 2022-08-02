package main

import (
	"io"
	"log"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {

	log.Println("Max was involked")

	var max uint32 = 0

	for {

		request, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Request received: %v", request.Number)

		if request.Number > max {
			max = request.Number

			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error while send response to client: %v\n", err)
			}
		}
	}

}

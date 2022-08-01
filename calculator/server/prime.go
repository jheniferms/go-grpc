package main

import (
	"fmt"

	pb "github.com/jheniferms/go-grpc/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	fmt.Printf("Primes was involked with: %v\n", in)
	var k uint32
	k = 2
	n := in.Number

	for n > 1 {
		if n%k == 0 {
			fmt.Println(k)
			n = n / k
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
		} else {
			k++
		}
	}
	return nil
}

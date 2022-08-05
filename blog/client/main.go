package main

import (
	"log"

	pb "github.com/jheniferms/go-grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// id := createBlog(c)
	// readBlog(c, id)
	// readBlog(c, "invalid")

	id := createBlog(c)
	readBlog(c, id)

	updateBook(c, id)
	readBlog(c, id)

}

package main

import (
	"context"
	"log"

	pb "github.com/jheniferms/go-grpc/blog/proto"
)

func updateBook(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Jheni Jheni",
		Title:    "trying to update",
		Content:  "Hiii",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Printf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}

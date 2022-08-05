package main

import (
	"context"
	"log"

	pb "github.com/jheniferms/go-grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "jheni",
		Title:    "my first book",
		Content:  "helloo",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}

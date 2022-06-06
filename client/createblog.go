package main

import (
	"context"
	"fmt"
	"log"
	pb "task1/protofile"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("------createBlog was invoked------")
	blog := &pb.Blog{
		Id:       "1",
		AuthorId: "100",
		Title:    "this is a title for Sha",
		Content:  "Content of the first blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		fmt.Printf("error Unexpected : %v\n", err)
	}
	fmt.Println(res.Id)
	return res.Id
}

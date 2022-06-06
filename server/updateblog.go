package main

import (
	"context"
	"fmt"
	"log"
	pb "task1/protofile"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*empty.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", in)
	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	if err := Session.Query("update allblog set authorid=?,title=?,content=? where id=?", data.AuthorId, data.Title, data.Content, in.Id).Exec(); err != nil {
		fmt.Println("error while update the Blog")
		fmt.Println(err)
	}
	return &emptypb.Empty{}, nil
}

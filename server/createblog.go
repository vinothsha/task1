package main

import (
	"context"
	"fmt"
	pb "task1/protofile"
)

func (s *Server) CreateBlog(ctx context.Context, req *pb.Blog) (*pb.BlogId, error) {
	// fmt.Println(req)
	data := BlogItem{
		Id:       req.Id,
		AuthorId: req.AuthorId,
		Content:  req.Content,
		Title:    req.Title,
	}
	if err := Session.Query("insert into allblog(id,authorid,title,content)values(?,?,?,?)", data.Id, data.AuthorId, data.Title, data.Content).Exec(); err != nil {
		fmt.Println("error while insert data into allblog")
		fmt.Println(err)
	}
	return &pb.BlogId{
		Id: req.Id,
	}, nil
}

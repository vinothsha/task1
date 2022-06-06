package main

import (
	"context"
	"fmt"
	pb "task1/protofile"
)

func (s *Server) ReadBlog(ctx context.Context, req *pb.BlogId) (*pb.Blog, error) {
	var data BlogItem
	m := map[string]interface{}{}
	iter := Session.Query("select * from allblog where id=?", req.Id).Iter()
	for iter.MapScan(m) {
		data.Id = m["id"].(string)
		data.AuthorId = m["authorid"].(string)
		data.Title = m["title"].(string)
		data.Content = m["content"].(string)
	}
	fmt.Println(data)
	return jsonToBlog(&data), nil
}

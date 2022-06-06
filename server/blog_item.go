package main

import (
	pb "task1/protofile"
)

type BlogItem struct {
	Id       string `json:"id"`
	AuthorId string `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func jsonToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.Id,
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}

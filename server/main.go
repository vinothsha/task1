package main

import (
	"fmt"
	"log"
	"net"
	pb "task1/protofile"

	"github.com/gocql/gocql"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var Session *gocql.Session

type Server struct {
	pb.BlogServiceServer
}

func main() {
	//db connect
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "blog"
	cluster.Consistency = gocql.LocalOne
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("error while connect to cassandra")
	} else {
		fmt.Println("cassandra is connected")
	}
	//server connect
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed tp listen on:%v\n", err)
	}
	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}

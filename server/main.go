package main

import (
	"context"
	"net"

	pb "github.com/yuki-toida/grpc-sample/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterTestServer(server, &Server{})
	server.Serve(listener)
}

type Server struct{}

func (s *Server) Get(c context.Context, r *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: r.Message}, nil
}

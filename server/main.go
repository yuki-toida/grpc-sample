package main

import (
	"context"
	"net"

	pb "github.com/yuki-toida/grpc-sample/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", port)
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

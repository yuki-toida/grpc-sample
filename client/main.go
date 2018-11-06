package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/yuki-toida/grpc-quick-start/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	message = "Hello"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	message := message
	if 1 < len(os.Args) {
		message = os.Args[1]
	}

	client := pb.NewTestClient(conn)
	res, err := client.Get(ctx, &pb.Request{Message: message})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

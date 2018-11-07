package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/yuki-toida/grpc-sample/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	message := "Hello"
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

package main

import (
	"time"
	"context"
	"fmt"
	"io"

	"github.com/wesley-lewis/learning-proto-buf/grpc_num"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := grpc_num.NewNumServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	stream, err := c.Rnd(ctx, &grpc_num.NumRequest{N: 5, From: 0, To: 100})
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				panic(err)
			}

			fmt.Println("Received:", resp.String())
		}
	}()
	<- done
}

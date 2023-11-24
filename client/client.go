package main 

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/wesley-lewis/learning-proto-buf/model"
	"time"
)

func main() {
	// Create new insecure credentials and block the client until the server is up 
	// if we don't use WithBlock opts the client doesn't block and the connection happens in the background
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()) , grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// We create new user service client and pass in the connection 
	c := model.NewUserServiceClient(conn)

	// We create a new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// We use the grpc Service which was registered in the .proto file, we might get a response too (depends on the definition)
	r, err := c.GetUser(ctx, &model.UserRequest{UserId: "Wesley"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Client received:", r.String())
}

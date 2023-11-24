package main

import (
	"context"
	"fmt"
	"net"

	"github.com/wesley-lewis/learning-proto-buf/model"
	"google.golang.org/grpc"
)

type UserServer struct {
	model.UnimplementedUserServiceServer
}

func( u *UserServer) GetUser(ctx context.Context, req *model.UserRequest) (*model.User, error) {
	fmt.Println("Server received:", req.String())
	return &model.User{UserId: "Wesley", Email: "wesley@gmail.com"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	model.RegisterUserServiceServer(s, &UserServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

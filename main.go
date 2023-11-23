// Title: GRPC

package main

import (
	"fmt"

	"context"

	"github.com/wesley-lewis/learning-proto-buf/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {
	userA := model.Group_User {
		UserId: "Wesley",
		Email: "wesley@gmail.com",
	}
	
	group := model.Group{
		Id: 1,
		Category: model.Category_DEVELOPER,
		Score: 23.3,
		Users: []*model.Group_User{&userA},
	}

	fmt.Println(&group)

	encoded, _ := proto.Marshal(&group)
	var grp model.Group
	proto.Unmarshal(encoded, &grp)
	fmt.Println(&grp)
}

type UserServiceClient interface {
	GetUser(ctx context.Context, in *model.UserRequest, opts grpc.CallOption) (*model.User, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func(c *userServiceClient) GetUser(ctx context.Context, in *model.UserRequest, opts grpc.CallOption) (*model.User, error) {
	return nil, nil
}

type UserServiceServer interface {
	GetUser(ctx context.Context, in *model.UserRequest, opts grpc.CallOption) (*model.User, error)
}

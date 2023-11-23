package main

import (
	"fmt"

	"github.com/wesley-lewis/learning-proto-buf/model"
	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("hello, world")
	user := model.User{
		UserId: "1",
		Email: "wesley@gmail.com",
	}
	user2 := model.User{
		UserId: "2",
		Email: "hrushi@gmail.com",
	}

	group := model.Group{
		Id: 1,
		Category: model.Category_DEVELOPER,
		Score: 2.4,
		Users: []*model.User{&user, &user2},
	}

	fmt.Println("To encode:", group.String())

	encoded, _ := proto.Marshal(&group)
	
	recovered := model.Group{}

	_ = proto.Unmarshal(encoded, &recovered)
	fmt.Println("Recovered:", recovered.String())
}

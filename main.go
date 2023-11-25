// Title: GRPC

package main

import (
	"fmt"


	"github.com/wesley-lewis/learning-proto-buf/model"
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


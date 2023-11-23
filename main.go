package main

import (
	"fmt"

	"github.com/wesley-lewis/learning-proto-buf/model"
	"google.golang.org/protobuf/proto"
)

func main() {
	group := model.Group{
		Id: 1,
		Category: model.Category_DEVELOPER,
		Score: 23.3,
	}

	fmt.Println(&group)

	encoded, _ := proto.Marshal(&group)
	var grp model.Group
	proto.Unmarshal(encoded, &grp)
	fmt.Println(&grp)
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wesley-lewis/learning-proto-buf/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Chat(stream chat.ChatService_SendTextClient, done chan bool) {
	t := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <- done:
			return 
		case <- t.C:
			err := stream.Send(&chat.ChatMessage{Txt: "Hello", Id: 1, To: 2})
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}

func Stats( stream chat.ChatService_SendTextClient, done chan bool) {
	for {
		stats, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(stats.String())
		if stats.TotalChar > 35 {
			fmt.Println("Beyond the limit")
			done <- true 
			stream.CloseSend()
			return 
		}
	}
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	stream, err := c.SendText(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go Stats(stream, done)
	go Chat(stream, done)

	<- done
}

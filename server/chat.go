package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/wesley-lewis/learning-proto-buf/chat"
	"google.golang.org/grpc"
)

type Server struct {
	chat.UnimplementedChatServiceServer
}

func(s *Server) SendText(stream chat.ChatService_SendTextServer) error {
	var total int64 = 0
	go func() {
		for {
			t := time.NewTicker(time.Second * 2)
			select {
				case <- t.C:
				stream.Send(&chat.ChatResponse{TotalChar: total})
			}
		}
	}()
	for {
		next, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Client closed")
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("->", next.Txt)
		total = total + int64(len(next.Txt))
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	handleErr(err)
	s := grpc.NewServer()
	chat.RegisterChatServiceServer(s, &Server{})
	err = s.Serve(lis)
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

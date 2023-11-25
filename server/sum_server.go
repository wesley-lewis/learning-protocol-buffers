package main

import (
	"io"
	"fmt"
	"net"
	"log"

	"github.com/wesley-lewis/learning-proto-buf/grpc_sum"
	"google.golang.org/grpc"
)

type Server struct {
	grpc_sum.UnimplementedNumServiceServer
}

func (s *Server) Sum(stream grpc_sum.NumService_SumServer) error {
	var total int64 = 0;
	var counter int = 0;
	for {
		next, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Received %d numbers sum: %d\n", counter, total)
			stream.SendAndClose(&grpc_sum.NumResponse{Total: total})
			return nil
		}
		if err != nil {
			return nil
		}

		total = total + next.X
		counter++
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpc_sum.RegisterNumServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

package main 

import (
	"net"
	"errors"
	"math/rand"
	"time"
	"fmt"
	
	"google.golang.org/grpc"
	"github.com/wesley-lewis/learning-proto-buf/grpc_num"
	
)

type NumServer struct {
	grpc_num.UnimplementedNumServiceServer
}

func(n *NumServer) Rnd(req *grpc_num.NumRequest, stream grpc_num.NumService_RndServer) error {
	fmt.Println(req.String())
	if req.N <= 0 {
		return errors.New("N must be greater than zero")
	}
	if req.To <= req.From {
		return errors.New("to must be greater or equal than from")
	}
	done := make(chan bool)

	go func() {
		for counter := 0; counter<int(req.N); counter++ {
			i := rand.Intn(int(req.To) - int(req.From) + 1) /*+ int(req.To)*/
			resp := grpc_num.NumResponse{I: int64(i), Remaining: req.N-int64(counter)}
			stream.Send(&resp)
			time.Sleep(time.Second)
		}
		done <- true
	}()

	<- done
	return nil 
}
func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	grpc_num.RegisterNumServiceServer(s, &NumServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

// package main
//
// import (
// 	"context"
// 	"time"
// 	"fmt"
// 	"log"
//
// 	"github.com/wesley-lewis/learning-proto-buf/grpc_sum"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )
//
// func main() {
// 	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()
//
// 	c := grpc_sum.NewNumServiceClient(conn)
//
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
// 	defer cancel()
//
// 	stream, err := c.Sum(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	from, to := 1, 100
// 	for i := from; i <= to; i++ {
// 		err := stream.Send(&grpc_sum.NumRequest{X: int64(i)})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
//
// 	fmt.Println("Waiting for response")
// 	result, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Printf("The sum form %d to %d is %d\n", from, to, result.Total)
// }
//
package main 

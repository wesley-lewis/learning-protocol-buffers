// package main
//
// import (
// 	"context"
// 	"fmt"
//
// 	"time"
//
// 	"github.com/wesley-lewis/learning-proto-buf/grpc_emp"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )
//
// // func main() {
// // 	// Create new insecure credentials and block the client until the server is up
// // 	// if we don't use WithBlock opts the client doesn't block and the connection happens in the background
// // 	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()) , grpc.WithBlock())
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	defer conn.Close()
// //
// // 	// We create new user service client and pass in the connection
// // 	c := model.NewUserServiceClient(conn)
// //
// // 	// We create a new context
// // 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// // 	defer cancel()
// //
// // 	// We use the grpc Service which was registered in the .proto file, we might get a response too (depends on the definition)
// // 	r, err := c.GetUser(ctx, &model.UserRequest{UserId: "Wesley"})
// // 	if err != nil {
// // 		panic(err)
// // 	}
// //
// // 	fmt.Println("Client received:", r.String())
// // }
//
// func main() {
// 	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
//
// 	c := grpc_emp.NewEmployeeInfoClient(conn)
//
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
//
// 	resp, err := c.GetAll(ctx, &grpc_emp.GetEmployeeRequest{})
// 	for _, emp := range resp.Employees {
// 		fmt.Println(emp.String())
// 	}
//
// 	// addr := &grpc_emp.Address {
// 	// 	Building: "Vibha",
// 	// 	Street: "Ambadi Road", 
// 	// 	Town: "Vasai", 
// 	// 	City: "Mumbai",
// 	// 	State: "Maharashtra",
// 	// }
// 	//
// 	// emp := &grpc_emp.SetEmployeeRequest{
// 	// 	Name: "Wesley Lewis", 
// 	// 	Address: addr,
// 	// }
// 	// r, err := c.Set(ctx, emp)
// 	// fmt.Println("Client received:", r.String())
// 	//
// 	// time.Sleep(time.Duration(3))
// 	// empReq := &grpc_emp.GetEmployeeRequest {
// 	// 	Id: r.Id,
// 	// }
// 	// resp, err := c.Get(ctx, empReq)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println("GET RPC:", resp.String())
// }

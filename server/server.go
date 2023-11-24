package main

import (
	"context"
	"fmt"
	"net"
	"strconv"

	grpc_emp "github.com/wesley-lewis/learning-proto-buf/grpc_emp"
	"google.golang.org/grpc"
)

// type UserServer struct {
// 	model.UnimplementedUserServiceServer
// }
//
// func( u *UserServer) GetUser(ctx context.Context, req *model.UserRequest) (*model.User, error) {
// 	fmt.Println("Server received:", req.String())
// 	return &model.User{UserId: "Wesley", Email: "wesley@gmail.com"}, nil
// }
//
// func main() {
// 	lis, err := net.Listen("tcp", "localhost:50051")
// 	if err != nil {
// 		panic(err)
// 	}
// 	s := grpc.NewServer()
// 	model.RegisterUserServiceServer(s, &UserServer{})
//
// 	if err := s.Serve(lis); err != nil {
// 		panic(err)
// 	}
// }

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	grpc_emp.RegisterEmployeeInfoServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}	
}

type Server struct {
	grpc_emp.UnimplementedEmployeeInfoServer
}

var employees = []*grpc_emp.Employee{}

func (s *Server) Get(ctx context.Context,req *grpc_emp.GetEmployeeRequest) (*grpc_emp.Employee, error) {
	id, _ := strconv.Atoi(req.Id)
	if len(employees) < id {
		return nil, fmt.Errorf("User not found")
	}
	return employees[id], nil
}

func (s *Server) Set(ctx context.Context, req *grpc_emp.SetEmployeeRequest) (*grpc_emp.SetEmployeeResponse, error) {
	id := fmt.Sprintf("{%d}", len(employees))
	emp := &grpc_emp.Employee {
		Name: req.Name,
		Id: id,
		Address: req.Address,
	}
	employees = append(employees, emp)
	return &grpc_emp.SetEmployeeResponse {
		Id: id,
	}, nil
}

func(s *Server) GetAll(context.Context, *grpc_emp.GetEmployeeRequest) (*grpc_emp.GetAllResponse, error) {
	return &grpc_emp.GetAllResponse{
		Employees: getEmployees(),
	}, nil	
}

func getEmployees() []*grpc_emp.Employee{
	var emps = make([]*grpc_emp.Employee, 1)
	for _, emp := range employees {
		emps = append(emps, emp)
	}
	return emps
}






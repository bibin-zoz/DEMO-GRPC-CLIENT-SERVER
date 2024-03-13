package main

import (
	"context"
	"log"
	"net"

	pb "go-grps-demo/proto/user"

	"google.golang.org/grpc"
)

type server struct {
	pb.UserServiceServer
}

func (s *server) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.User, error) {
	return &pb.User{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("gRPC server started on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

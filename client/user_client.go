package main

import (
	"context"
	pb "go-grps-demo/proto/user"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	userResp, err := client.GetUserById(context.Background(), &pb.GetUserByIdRequest{Id: 123})
	if err != nil {
		log.Fatalf("error getting user: %v", err)
	}

	log.Printf("User: %v", userResp)
}

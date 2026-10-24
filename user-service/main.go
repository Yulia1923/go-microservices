package main

import (
	"log"
	"net"

	"go-microservices/user-service/userpb"

	"google.golang.org/grpc"
)

func main() {
	users := map[int64]User{
		1001: {
			ID:    1001,
			Name:  "Alice",
			Email: "alice@example.com",
		},
		1002: {
			ID:    1002,
			Name:  "Bob",
			Email: "bob@example.com",
		},
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userServer := &UserServer{
		users: users,
	}

	userpb.RegisterUserServiceServer(grpcServer, userServer)

	log.Println("User Service listening on:50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

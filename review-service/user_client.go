package main

import (
	"context"
	"time"

	"go-microservices/user-service/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}
}

func (c *UserClient) GetUser(id int64) (*userpb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return c.client.GetUser(ctx, &userpb.GetUserRequest{
		Id: id,
	})
}
func DialUserService() (*UserClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	return NewUserClient(conn), conn, nil
}

package main

import (
	"context"

	"go-microservices/user-service/userpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	users map[int64]User
}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *userpb.GetUserRequest,
) (*userpb.GetUserResponse, error) {
	user, exists := s.users[req.GetId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &userpb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

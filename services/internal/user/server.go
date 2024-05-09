package userserver

import (
	"context"

	userservice "mariusihring.dev/services/rpc/user"
)

type UserServer struct {
	userservice.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, req *userservice.UserRequest) (*userservice.UserResponse, error) {
	return &userservice.UserResponse{
		Uid:         "1",
		DisplayName: "test",
		Phone:       "123213",
		Provider:    "test",
		Created:     "12.12.12",
		LastSignIn:  "12.12.12",
	}, nil
}

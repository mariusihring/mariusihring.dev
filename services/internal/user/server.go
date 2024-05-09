package userserver

import (
	"context"
	pb "mariusihring.dev/services/rpc/user"
)

type UserServer struct{}

func (s *UserServer) GetUser(ctx context.Context, reg *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{
		Uid:         "",
		DisplayName: "",
		Phone:       "",
		Provider:    "",
		Created:     "",
		LastSignIn:  "",
	}, nil
}

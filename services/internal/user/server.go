package userserver

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"time"

	userservice "mariusihring.dev/services/rpc/user"
)

type UserServer struct {
	userservice.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(ctx context.Context, req *userservice.UserRequest) (*userservice.UserResponse, error) {
	return &userservice.UserResponse{
		Id:          1,
		LastName:     "ihring",
		FirstName:    "marius",
		UserName:     "marius_ihring",
		Email:        "marius.ihring@icloud.com",
		PasswordHash: "akjsdhakj",
		CreatedAt:    "2024-04-04",
		UpdatedAt:    "2024-04-04",
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *userservice.NewUserRequest) (*userservice.UserResponse, error) {

	return &userservice.UserResponse{
		Id:          1,
		LastName:     req.LastName,
		FirstName:    req.FirstName,
		UserName:     req.UserName,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}, nil

}

func (s *UserServer) UpdateUser(ctx context.Context, req *userservice.NewUserRequest) (*userservice.UserResponse, error) {

	return &userservice.UserResponse{
		Id:          1,
		LastName:     req.LastName,
		FirstName:    req.FirstName,
		UserName:     req.UserName,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		CreatedAt:    "2024-04-04",
		UpdatedAt:    time.Now().String(),
	}, nil

}

func NewUserServiceServer(lc fx.Lifecycle) *grpc.Server {
	user_service_port := os.Getenv("USER_SERVICE_PORT")
	port := fmt.Sprintf(":%s", user_service_port)
	svr := grpc.NewServer()
	userservice.RegisterUserServiceServer(svr, &UserServer{})
	reflection.Register(svr)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", port)
			if err != nil {
				return err
			}
			fmt.Println("Starting UserService Server at ", port)
			go svr.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			svr.GracefulStop()
			return nil
		},
	})
	return svr
}

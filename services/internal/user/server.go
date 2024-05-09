package userserver

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"

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

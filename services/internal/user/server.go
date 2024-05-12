package userserver

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"mariusihring.dev/services/db/databasemodels"
	"net"
	"os"
	"time"

	userservice "mariusihring.dev/services/rpc/user"
)

type UserServer struct {
	userservice.UnimplementedUserServiceServer
	DB *gorm.DB
}

func (s *UserServer) GetUser(ctx context.Context, req *userservice.UserRequest) (*userservice.UserResponse, error) {
	user := databasemodels.User{Id: uint(req.Id)}
	result := s.DB.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &userservice.UserResponse{
		Id:           int64(user.Id),
		LastName:     *user.LastName,
		FirstName:    *user.FirstName,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.String(),
		UpdatedAt:    user.UpdatedAt.String(),
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *userservice.NewUserRequest) (*userservice.UserResponse, error) {
	user := databasemodels.User{
		LastName:     &req.LastName,
		FirstName:    &req.FirstName,
		UserName:     req.UserName,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		Session:      nil,
		Goals:        nil,
	}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userservice.UserResponse{
		Id:           1,
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
	user := databasemodels.User{Email: req.Email}
	result := s.DB.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	user.LastName = &req.LastName
	user.FirstName = &req.FirstName
	user.UserName = req.UserName
	user.Email = req.Email
	user.PasswordHash = req.PasswordHash

	s.DB.Save(&user)
	return &userservice.UserResponse{
		Id:           int64(user.Id),
		LastName:     *user.LastName,
		FirstName:    *user.FirstName,
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    "2024-04-04",
		UpdatedAt:    time.Now().String(),
	}, nil

}

func NewUserServiceServer(lc fx.Lifecycle, db *gorm.DB) *grpc.Server {
	user_service_port := os.Getenv("USER_SERVICE_PORT")
	port := fmt.Sprintf(":%s", user_service_port)
	svr := grpc.NewServer()
	userservice.RegisterUserServiceServer(svr, &UserServer{DB: db})
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
	return nil
}

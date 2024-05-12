package clients

import (
	"google.golang.org/grpc"
	"log"

	userpb "mariusihring.dev/services/rpc/user"
)

type UserServiceClient struct {
	Client userpb.UserServiceClient
}

func NewUserServiceClient(addr string) *UserServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("UserServiceClient did not connect: %v", err)
	}
	client := userpb.NewUserServiceClient(conn)
	return &UserServiceClient{Client: client}
}
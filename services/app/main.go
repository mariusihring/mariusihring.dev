package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	userserver "mariusihring.dev/services/internal/user"
	user "mariusihring.dev/services/rpc/user"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error laoding .env")
	}
	user_service_port := os.Getenv("USER_SERVICE_PORT")
	port := fmt.Sprintf(":%s", user_service_port)
	println(port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &userserver.UserServer{})
	reflection.Register(server)
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

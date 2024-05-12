package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"mariusihring.dev/services/db"
	userserver "mariusihring.dev/services/internal/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fx.New(
		fx.Provide(
			db.NewDbConnection,
			userserver.NewUserServiceServer,
			),
		fx.Invoke(func(*grpc.Server) {}),
	).Run()
}

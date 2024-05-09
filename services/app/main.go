package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"mariusihring.dev/services/internal/user"
	"mariusihring.dev/services/rpc/user"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := &userserver.UserServer{}
	userHandler := user.NewUserServiceServer(server)

	port := os.Getenv("USER_SERVICE_PORT") // Replace "PORT" with your environment variable name

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	addr := fmt.Sprintf(":%s", port)

	http.ListenAndServe(addr, userHandler)
}

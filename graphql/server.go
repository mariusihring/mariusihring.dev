package main

import (
	"github.com/joho/godotenv"
	"log"
	"mariusihring.dev/graphql/clients"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"mariusihring.dev/graphql/graph"
	resolvers "mariusihring.dev/graphql/graph/resolvers"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	user_service_port := os.Getenv("USER_SERVICE_PORT")

	userServiceClient := clients.NewUserServiceClient(user_service_port)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		UserService: userServiceClient,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

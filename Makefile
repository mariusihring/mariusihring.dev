start_backend:
	make start_services & make start_graphql



start_services:
	go run services/app/main.go

start_graphql:
	go run graphql/server.go
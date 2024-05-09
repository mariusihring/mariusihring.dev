package databaseserver

import (
	databaseservice "mariusihring.dev/services/rpc/database"
)

type DatabaseServer struct {
	databaseservice.UnimplementedDatabaseServiceServer
}

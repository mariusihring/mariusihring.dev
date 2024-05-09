package valorantserver

import (
	valorantservice "mariusihring.dev/services/rpc/valorant"
)

type ValorantServer struct {
	valorantservice.UnimplementedValorantServiceServer
}

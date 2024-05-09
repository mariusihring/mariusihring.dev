package savingserver

import (
	savingservice "mariusihring.dev/services/rpc/saving"
)

type SavingServer struct {
	savingservice.UnimplementedSavingServiceServer
}

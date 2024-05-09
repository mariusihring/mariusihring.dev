package configserver

import configservice "mariusihring.dev/services/rpc/config"

type ConfigServer struct {
	configservice.UnimplementedConfigServiceServer
}

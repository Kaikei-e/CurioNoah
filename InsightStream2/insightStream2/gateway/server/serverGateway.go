package server

import "insightStream2/driver/rest"

type Gateway interface {
	Initialize()
	HealthCheck() error
}

type GatewayImpl struct{}

func NewServerGateway() Gateway {
	return &GatewayImpl{}
}

func (s *GatewayImpl) Initialize() {
	rest.Initialize()
}

func (s *GatewayImpl) HealthCheck() error {
	panic("implement me")
}

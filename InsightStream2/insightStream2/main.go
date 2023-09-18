package main

import "insightStream2/gateway/server"

func main() {
	server := server.GatewayImpl{}

	err := server.Initialize()
	if err != nil {
		panic(err)
	}
	err = server.HealthCheck()
	if err != nil {
		panic(err)
	}
}

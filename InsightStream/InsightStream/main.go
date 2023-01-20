package main

import (
	"insightstream/repository"
	"insightstream/server"
)

func main() {
	cl := repository.InitConnection()

	server.Server(cl)

}

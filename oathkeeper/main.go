package main

import (
	"ghostkeeper/logger"
	"ghostkeeper/rest"
)

func main() {
	logger.InitLogger()

	rest.Server()
}

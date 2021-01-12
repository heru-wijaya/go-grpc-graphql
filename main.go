package main

import (
	logger "github.com/heru-wijaya/go-grpc-skeleton/lib"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger.Log("test", log.Fields{
		"test": "testing",
	}, "error")
}

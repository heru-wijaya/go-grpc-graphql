package main

import (
	logger "github.com/heru-wijaya/go-grpc-skeleton/lib"
)

func main() {
	name := "first last"

	logger.Log("test", map[string]interface{}{
		"name": name,
	}, "error")
}

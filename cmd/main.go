package main

import (
	"user/grpc"

	"github.com/go-gin/gin"
)

func main() {
	//

	router := gin.Default()
	go grpc.RunGRPC()

	router.Run(":9090")
}

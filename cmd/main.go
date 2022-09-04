package main

import (
	"github.com/shaineminkyaw/RabbitTest/rpc"

	"github.com/go-gin/gin"
)

func main() {
	//

	router := gin.Default()
	go rpc.RunGRPC()

	router.Run(":9090")
}

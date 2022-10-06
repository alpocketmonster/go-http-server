package main

import (
	"log"
	"os"

	"github.com/alpocketmonster/GoHttpServer/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin.Default()

	controller := new(controller.Controller)

	r.GET("/auth", controller.AuthMiddleware)

	// Listen and Server in https://127.0.0.1:8080
	r.Run(":8080")
}

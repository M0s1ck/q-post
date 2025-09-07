package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"user-service/internal/delivery"
)

func main() {
	engine := gin.Default()
	engine.GET("/health", delivery.HealthCheck)

	addr := ":8080"
	err := engine.Run(addr)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Hello, world!")
}

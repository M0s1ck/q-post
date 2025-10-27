package main

import (
	"fmt"
	"log"
	"os"
	"user-service/internal/infra/env"

	"user-service/internal/app"
)

// Swagger attributes:
//
// @title QPost user service
// @version 1.0.0
// @description Gin app to deal with users
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	envConf := env.BuildEnvConfig()

	engine := app.BuildGinEngine(envConf)

	addr := ":" + os.Getenv(envConf.AppPort)
	err := engine.Run(addr)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Hello, world!")
}

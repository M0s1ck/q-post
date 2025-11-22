package main

import (
	"log"

	"auth-service/internal/app"
	"auth-service/internal/infra/env"
)

// Swagger attributes:
//
// @title QPost auth service
// @version 1.0.0
// @description Gin app for user auth
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	log.Println("Hello, auth!")
	envConf := env.BuildEnvConfig()

	engine := app.BuildGinEngine(envConf)

	addr := ":" + envConf.AppPort
	err := engine.Run(addr)

	if err != nil {
		log.Println("Gin engine start err: ", err)
		panic(err)
	}
}

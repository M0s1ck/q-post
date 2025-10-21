package main

import (
	"log"
	"os"

	"auth-service/internal/app"
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
	app.LoadEnv()

	engine := app.BuildGinEngine()

	addr := ":" + os.Getenv("AUTH_SERVICE_PORT")
	err := engine.Run(addr)

	if err != nil {
		log.Println("Gin engine start err: ", err)
		panic(err)
	}
}

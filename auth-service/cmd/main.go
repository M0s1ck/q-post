package main

import (
	"auth-service/internal/app"
	"fmt"
)

// Swagger attributes:
//
// @title QPost auth service
// @version 1.0.0
// @description Gin app for user auth
// @schemes http https
func main() {
	fmt.Println("Hello, auth!")

	engine := app.BuildGinEngine()

	addr := ":8088"
	err := engine.Run(addr)

	if err != nil {
		fmt.Println("Gin engine start err: ", err)
		panic(err)
	}
}

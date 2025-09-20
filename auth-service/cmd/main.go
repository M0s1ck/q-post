package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"auth-service/internal/app"
	"auth-service/internal/infra/db"
)

// Swagger attributes:
//
// @title QPost auth service
// @version 1.0.0
// @description Gin app for user auth
// @schemes http https
func main() {
	var psg *gorm.DB = db.ConnectToPostgres()
	log.Println(*psg)

	fmt.Println("Hello, auth!")

	engine := app.BuildGinEngine()

	addr := ":8088"
	err := engine.Run(addr)

	if err != nil {
		fmt.Println("Gin engine start err: ", err)
		panic(err)
	}
}

package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"user-service/internal/app"
	"user-service/internal/infra/db"
)

// Swagger attributes:
//
// @title QPost user service
// @version 1.0.0
// @description Gin app to deal with users
// @schemes http https
func main() {
	var psg *gorm.DB = db.ConnectToPostgres()
	log.Println(psg)

	engine := app.BuildGinEngine(psg)

	addr := ":8080"
	err := engine.Run(addr)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Hello, world!")
}

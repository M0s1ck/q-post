package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"

	"user-service/internal/app"
	"user-service/internal/infra/db"
)

func main() {
	var psg *gorm.DB = db.ConnectToPostgres()
	log.Println(psg)

	engine := app.Build()

	addr := ":8080"
	err := engine.Run(addr)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Hello, world!")
}

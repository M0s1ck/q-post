package main

import (
	"fmt"
	"log"
	"user-service/internal/infra/db"

	"user-service/internal/app"
)

func main() {
	engine := app.Build()

	db.ConnectToPostgres()

	addr := ":8080"
	err := engine.Run(addr)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Hello, world!")
}

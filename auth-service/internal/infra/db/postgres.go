package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

const dbConnectionStr = "postgres://postgres:postgres@localhost:5432/q-post?sslmode=disable&search_path=auth"

func ConnectToPostgres() *gorm.DB {
	var psgConf postgres.Config = postgres.Config{
		DSN:                  dbConnectionStr,
		PreferSimpleProtocol: true,
	}

	var dialector gorm.Dialector = postgres.New(psgConf)

	db, err := gorm.Open(dialector) // TODO: here pointer to config for some reason

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	sqlDB, _ := db.DB()
	fmt.Println(sqlDB.Stats())

	return db
}

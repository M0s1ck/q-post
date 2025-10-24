package infra

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres() *gorm.DB {
	dbConnectionStr := getDbConnectionString()

	var psgConf postgres.Config = postgres.Config{
		DSN:                  dbConnectionStr,
		PreferSimpleProtocol: true,
	}

	var dialector gorm.Dialector = postgres.New(psgConf)

	db, err := gorm.Open(dialector)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	sqlDB, _ := db.DB()
	fmt.Println(sqlDB.Stats())

	return db
}

func getDbConnectionString() string {
	var host string
	if os.Getenv("IN_DOCKER") == "1" {
		host = os.Getenv("POSTGRES_HOST")
	} else {
		host = "localhost"
	}

	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")),
		Host:   host + ":" + os.Getenv("POSTGRES_PORT"),
		Path:   os.Getenv("POSTGRES_DB"),
	}

	q := u.Query()
	q.Set("sslmode", "disable")
	q.Set("search_path", os.Getenv("POSTGRES_AUTH_SCHEME"))

	u.RawQuery = q.Encode()
	log.Println(u.String())
	return u.String() // smth like postgres://postgres:postgres@localhost:5432/q-post?sslmode=disable&search_path=auth"
}

// postgres://postgres:postgres@:5432/q-post?search_path=auth&sslmode=disable
// postgres://postgres:postgres@psg:5432/q-post?sslmode=disable

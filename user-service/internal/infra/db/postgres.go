package db

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"user-service/internal/infra/env"
)

func ConnectToPostgres(conf *env.PostgresConfig) *gorm.DB {
	dbConnectionStr := getDbConnectionString(conf)

	var psgConf postgres.Config = postgres.Config{
		DSN:                  dbConnectionStr,
		PreferSimpleProtocol: true,
	}

	var dialector gorm.Dialector = postgres.New(psgConf)

	db, err := gorm.Open(dialector)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	sqlDB, _ := db.DB()
	fmt.Println(sqlDB.Stats())

	return db
}

func getDbConnectionString(conf *env.PostgresConfig) string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.User, conf.Password),
		Host:   conf.Host + ":" + conf.Port,
		Path:   conf.DB,
	}

	q := u.Query()
	q.Set("sslmode", "disable")
	q.Set("search_path", conf.Scheme)

	u.RawQuery = q.Encode()
	log.Println(u.String())
	return u.String() // smth like postgres://postgres:postgres@localhost:5432/q-post?sslmode=disable&search_path=community"
}

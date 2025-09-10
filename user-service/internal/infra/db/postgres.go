package db

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"user-service/internal/domain"
)

const dbConnectionStr = "postgres://postgres:postgres@localhost:5432/q-post?sslmode=disable"

func ConnectToPostgres() *gorm.DB {
	ctx := context.Background()
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

	// Temp
	user := domain.User{
		Id:       uuid.New(),
		Username: "dummy2",
	}

	err = gorm.G[domain.User](db).Create(ctx, &user)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create user: %v\n", err)
		os.Exit(1)
	}

	return db
}

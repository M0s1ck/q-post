package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func ConnectToPostgres() {
	ctx := context.Background() // "postgres://username:password@localhost:5432/database_name"
	dbpool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/q-post")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	pingErr := dbpool.Ping(ctx)
	if pingErr != nil {
		log.Printf("Unable to ping database: %v\n", pingErr)
	}
}

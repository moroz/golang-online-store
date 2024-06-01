package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/moroz/sqlc-demo/handler"
)

func MustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}
	return value
}

var DatabaseUrl = MustGetenv("DATABASE_URL")

func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, DatabaseUrl)
	defer db.Close(ctx)

	if err != nil {
		log.Fatal(err)
	}

	router := handler.Router(db)
	router.Run(":3000")
}

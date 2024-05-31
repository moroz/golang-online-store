package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/moroz/sqlc-demo/db/queries"
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
	conn, err := pgx.Connect(ctx, DatabaseUrl)
	defer conn.Close(ctx)

	if err != nil {
		log.Fatal(err)
	}

	queries := queries.New(conn)

	products, err := queries.ListProducts(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", products)
}

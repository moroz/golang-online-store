package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/moroz/sqlc-demo/controllers"
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

	r := gin.Default()

	products := controllers.ProductController(conn)
	r.GET("/", products.Index)

	r.Run(":3000")
}

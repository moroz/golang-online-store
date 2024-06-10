package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/moroz/sqlc-demo/config"
	"github.com/moroz/sqlc-demo/handler"
)

func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, config.DatabaseUrl)
	defer db.Close(ctx)

	if err != nil {
		log.Fatal(err)
	}

	router := handler.Router(db)
	router.Run(":3000")
}

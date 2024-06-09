package main

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"github.com/gin-contrib/sessions/cookie"
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

func MustGetenvBase64(key string) []byte {
	str := MustGetenv(key)
	binary, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalf("FATAL: Failed to decode environment variable %s as Base64-encoded binary!", key)
	}
	return binary
}

var DatabaseUrl = MustGetenv("DATABASE_URL")
var SessionSigner = MustGetenvBase64("SESSION_SIGNER_BASE64")

func main() {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, DatabaseUrl)
	defer db.Close(ctx)

	if err != nil {
		log.Fatal(err)
	}

	cookieStore := cookie.NewStore(SessionSigner)
	router := handler.Router(db, cookieStore)
	router.Run(":3000")
}

package main

import (
	"log"
	"os"
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

}

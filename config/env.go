package config

import (
	"encoding/base64"
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

package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}

// use godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

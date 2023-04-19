package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("MONGO_URI")
}
func EnvMongoDBName() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("DB_NAME")
}

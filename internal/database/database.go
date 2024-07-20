package database

import (
	"log"

	"github.com/joho/godotenv"
)

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

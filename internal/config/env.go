package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() error {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
	return nil
}
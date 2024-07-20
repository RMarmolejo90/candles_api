package database

import (
	"log"
	"os"

	"github.com/rmarmolejo90/candles_api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	config.LoadEnv()
	dsn := os.Getenv("DB_DSN")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Printf("\n\n*** Connected To The Database ***\n\n")

	// Need to add all models for migration
	err = DB.AutoMigrate()
	if err != nil {
		log.Print("Error Migrating the database: " + err.Error())
	}
}

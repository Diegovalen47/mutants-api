package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {

	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Println("Error loading .env file")
	}

	var dsn string

	env := os.Getenv("IS_LOCAL")

	if env == "true" {
		dsn = "host=" + os.Getenv("HOST") + " port=" + os.Getenv("PORT") + " user=" + os.Getenv("USER_DB") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable"
	} else {
		dsn = "host=" + os.Getenv("HOST") + " user=" + os.Getenv("USER_DB") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=require"
	}

	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal("Error connecting to database: ", error)
	} else {
		log.Println("Connected to database")
	}
}

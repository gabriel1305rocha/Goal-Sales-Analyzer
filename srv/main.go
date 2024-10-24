package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gabriel1305rocha/Goal-Sales-Analyzer/controllers"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSLMODE")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword +
		" dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSslMode

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	http.HandleFunc("/HelloWorld", func(w http.ResponseWriter, r *http.Request) {
		controller.HelloWorld(db, w, r)
	})

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

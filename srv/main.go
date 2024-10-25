package main

import (
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/gabriel1305rocha/Goal-Sales-Analyzer/models"
	"github.com/gabriel1305rocha/Goal-Sales-Analyzer/routers"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
}

func loadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	config := &Config{
		Host:     getEnv("DB_HOST"),
		User:     getEnv("DB_USER"),
		Password: getEnv("DB_PASSWORD"),
		DbName:   getEnv("DB_NAME"),
		Port:     getEnv("DB_PORT"),
		SslMode:  getEnv("DB_SSLMODE"),
	}

	return config, nil
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("environment variable %s is not set", key)
	}
	return value
}

func initDB() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DbName, config.Port, config.SslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	models.Db = db
	if err := db.AutoMigrate(&models.User{}, &models.Sales{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func main() {
	initDB()
	routers.Init()
	beego.Run()
}

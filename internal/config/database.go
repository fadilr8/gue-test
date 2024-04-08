package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fadilr8/gue-test/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	println("Connected to the database")
	
	database, err := gorm.Open(mysql.Open(connectionString))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&model.User{}, &model.Employee{})

	DB = database
}

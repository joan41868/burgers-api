package main

import (
	"burger-api/domain/application"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("dbHost"),
		os.Getenv("dbUsername"), os.Getenv("dbPassword"), os.Getenv("dbName"), os.Getenv("dbPort"))
	app := application.NewApplication(connStr)
	app.Start(os.Getenv("PORT"))
}

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codeInBit/mkobo-test/controllers"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	var server = controllers.Server{}

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file : %v", err)
	} else {
		fmt.Println("Success loading .env file")
	}

	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	port := os.Getenv("PORT")
	server.Run(":" + port)
}

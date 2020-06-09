package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codeInBit/mkobo-test/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file :")
	} else {
		fmt.Printf("Successfully loaded environmental variable\n")
	}

	port := os.Getenv("PORT")

	// Handle routes
	http.Handle("/", routes.Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

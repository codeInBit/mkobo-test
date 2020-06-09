package utilities

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect interface
	"github.com/joho/godotenv"
)

//ConnectDB - Make database connection
func ConnectDB() *gorm.DB {

	//Load environmenatal variables
	err := godotenv.Load()

	var db *gorm.DB

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	//Define DB connection string
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	//connect to db URI
	db, err = gorm.Open("mysql", dbURI)

	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	} else {
		fmt.Printf("We are connected to the mysql database\n")
	}
	// close db when not in use
	defer db.Close()

	// Migrate the schema
	db.Debug().AutoMigrate()

	return db
}

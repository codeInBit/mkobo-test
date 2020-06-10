package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/codeInBit/mkobo-test/models"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
)

//Server - A struct that defines DB field to initialize db connection and Route to load route file
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

//Initialize - Initialize connection to mysql database
func (s *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	s.DB, err = gorm.Open("mysql", DbURI)
	if err != nil {
		fmt.Printf("Cannot connect to mysql database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("Successfully connected to database\n")
	}

	//database migration
	s.DB.Debug().AutoMigrate(&models.User{}, &models.PasswordReset{})

	s.Router = mux.NewRouter()
	s.LoadRoutes()
}

//Run - Serve the application
func (s *Server) Run(addr string) {
	fmt.Println("Listening to port ", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

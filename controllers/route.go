package controllers

import (
	"github.com/codeInBit/mkobo-test/middlewares"
)

//LoadRoutes - This is where all the routes are defined
func (s *Server) LoadRoutes() {
	api := s.Router.PathPrefix("/api").Subrouter()
	//Home Route
	api.HandleFunc("", middlewares.CommonMiddleware(s.Home)).Methods("GET")

	//User routes
	api.HandleFunc("/users/register", middlewares.CommonMiddleware(s.Register)).Methods("POST")
}

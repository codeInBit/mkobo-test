package controllers

import (
	"github.com/codeInBit/mkobo-test/middlewares"
)

//LoadRoutes -
func (s *Server) LoadRoutes() {
	//Home Route
	s.Router.HandleFunc("/", middlewares.CommonMiddleware(s.Home)).Methods("GET")

}

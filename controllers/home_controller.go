package controllers

import (
	"net/http"

	"github.com/codeInBit/mkobo-test/utilities"
)

//Home - This handles the "/" endpoint just to display a welcome message
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	utilities.JSON(w, http.StatusOK, "Welcome, MKOBO", "Successful")
}

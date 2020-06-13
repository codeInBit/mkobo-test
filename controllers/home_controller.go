package controllers

import (
	"net/http"

	"github.com/codeInBit/wallet-app/utilities"
)

//Home - This handles the "/" endpoint just to display a welcome message
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	utilities.JSON(w, http.StatusOK, "Welcome, MKOBO \n This application contains the endpoints built for MKOBO Test. \n The documentation for the API is accessible at 'https://documenter.getpostman.com/view/947913/SzzehgDD'", "Successful")
}

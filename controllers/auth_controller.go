package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codeInBit/mkobo-test/models"
	"github.com/codeInBit/mkobo-test/utilities"
)

//Register - This method registers user
func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("register")
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(s.DB)

	if err != nil {

		formattedError := utilities.FormatError(err.Error())
		utilities.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	utilities.JSON(w, http.StatusCreated, userCreated)
}

//Login - This method logs users in
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := user.SignIn(user.Email, user.Password, s.DB)
	if err != nil {
		formattedError := utilities.FormatError(err.Error())
		utilities.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	
	utilities.JSON(w, http.StatusOK, token)
}

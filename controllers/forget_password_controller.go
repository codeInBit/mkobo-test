package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/codeInBit/mkobo-test/models"
	"github.com/codeInBit/mkobo-test/utilities"
)

//ForgotPassword - This method accepts sends email containing link to reset a user's password
func (s *Server) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Sorry, An error occured!")
		return
	}
	user := models.User{}
	passwordreset := models.PasswordReset{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Sorry, An error occured!")
		return
	}

	user.Prepare()
	err = user.Validate("forgotpassword")
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Validation error, please check all required fields")
		return
	}
	result, err := user.FindUserByEmail(user.Email, s.DB)
	if err != nil {
		formattedError := utilities.FormatError(err.Error())
		utilities.ERROR(w, http.StatusUnprocessableEntity, formattedError, "Error")
		return
	}

	//Delete existing reset token record and a new Save password reset token
	passwordreset.Email = user.Email
	_, err = passwordreset.DeleteAResetRecord(user.Email, s.DB)
	userResetToken, err := passwordreset.SaveResetToken(s.DB)

	//Send email
	utilities.SendEmail(*userResetToken)

	utilities.JSON(w, http.StatusOK, result, "To reset your password, follow the details sent to your email address.")
}

//ResetLink - This method accepts sends email containing link to reset a user's password
func (s *Server) ResetLink(w http.ResponseWriter, r *http.Request) {
}

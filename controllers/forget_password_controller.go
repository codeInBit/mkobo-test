package controllers

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/codeInBit/wallet-app/models"
	"github.com/codeInBit/wallet-app/utilities"
)

//ForgotPassword - This method accepts user email, and sends a mail containing link to reset a user's password
func (s *Server) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Sorry, An error occured!")
		return
	}

	//Assign needed models
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
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	//Fetch user using email
	result, err := user.FindUserByEmail(user.Email, s.DB)
	if err != nil {
		formattedError := utilities.FormatError(err.Error())
		utilities.ERROR(w, http.StatusUnprocessableEntity, formattedError, "")
		return
	}
	result.Password = ""

	//Delete existing reset token record and a new Save password reset token
	passwordreset.Email = user.Email
	_, err = passwordreset.DeleteAResetRecord(user.Email, s.DB)
	userResetToken, err := passwordreset.SaveResetToken(s.DB)

	//Send email
	utilities.SendEmail(*userResetToken)

	utilities.JSON(w, http.StatusOK, result, "To reset your password, follow the details sent to your email address.")
}

//ResetPassword - This method accepts the reset token, and new password to be updated
func (s *Server) ResetPassword(w http.ResponseWriter, r *http.Request) {
	//Struct that defines input values for this request so we can bind it
	type ResetData struct {
		Token                string `json:"token"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"passwordConfirmation"`
	}

	//Assigned needed models
	resetData := ResetData{}
	passwordreset := models.PasswordReset{}
	user := models.User{}

	var updatedUser *models.User

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Sorry, An error occured!")
		return
	}

	err = json.Unmarshal(body, &resetData)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Sorry, An error occured!")
		return
	}

	//Perform validation on TransferData inputs
	resetData.Password = html.EscapeString(strings.TrimSpace(resetData.Password))
	resetData.PasswordConfirmation = html.EscapeString(strings.TrimSpace(resetData.PasswordConfirmation))

	if resetData.Token == "" {
		utilities.ERROR(w, http.StatusUnauthorized, errors.New("Reset token is Required"), "")
		return
	}
	if resetData.Password == "" {
		utilities.ERROR(w, http.StatusUnauthorized, errors.New("Password is Required"), "")
		return
	}
	if resetData.PasswordConfirmation == "" {
		utilities.ERROR(w, http.StatusUnauthorized, errors.New("Confirm password is Required"), "")
		return
	}

	//Confirm if token exist
	resetRecord, err := passwordreset.FindResetRecordByToken(resetData.Token, s.DB)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	//Update Password of the user assigned to the token
	if resetData.Password == resetData.PasswordConfirmation {
		updatedUser, err = user.ChangePassword(resetRecord.Email, resetData.Password, s.DB)
		updatedUser.Password = ""
		if err != nil {
			utilities.ERROR(w, http.StatusInternalServerError, err, "")
			return
		}

		//Delete token record from password reset table
		_, err = passwordreset.DeleteAResetRecord(updatedUser.Email, s.DB)

		utilities.JSON(w, http.StatusOK, updatedUser, "Password Updated Successfully")
		return
	} else {
		utilities.ERROR(w, http.StatusUnprocessableEntity, nil, "Passwords do not match")
		return
	}
}

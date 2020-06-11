package controllers

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/codeInBit/mkobo-test/auth"
	"github.com/codeInBit/mkobo-test/models"
	"github.com/codeInBit/mkobo-test/utilities"
)

//Transfer - This method allow tranfer of value between two users
func (s *Server) Transfer(w http.ResponseWriter, r *http.Request) {
	//Struct that defines input values for this request so we can bind it
	type TransferData struct {
		Amount int    `json:"amount"`
		Email  string `json:"email"`
	}

	//Assigned needed models
	transferData := TransferData{}
	user := models.User{}
	wallet := models.Wallet{}
	// wallettransaction := models.WalletTransaction{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Sorry, An error occured!")
	}

	err = json.Unmarshal(body, &transferData)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "Sorry, An error occured!")
		return
	}

	//Perform validation on TransferData inputs
	transferData.Email = html.EscapeString(strings.TrimSpace(transferData.Email))

	if transferData.Amount == 0 {
		utilities.ERROR(w, http.StatusUnauthorized, errors.New("Amount Required"), "")
		return
	}
	if transferData.Email == "" {
		utilities.ERROR(w, http.StatusUnauthorized, errors.New("Email Required"), "")
		return
	}

	//Get Logged in UserID
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		utilities.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"), "Invalid Token")
		return
	}

	//Get Recipient's User Object
	recipientUser, err := user.FindUserByEmail(transferData.Email, s.DB)
	if err != nil {
		utilities.ERROR(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	//Fund Wallet
	err = wallet.Transfer(userID, recipientUser.ID, transferData.Amount, s.DB)
	if err != nil {
		utilities.ERROR(w, http.StatusBadRequest, err, "")
		return
	}

	utilities.JSON(w, http.StatusOK, err, "Transfer was successfull")
}

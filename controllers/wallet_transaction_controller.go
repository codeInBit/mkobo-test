package controllers

import (
	"errors"
	"net/http"

	"github.com/codeInBit/mkobo-test/auth"
	"github.com/codeInBit/mkobo-test/models"
	"github.com/codeInBit/mkobo-test/utilities"
)

//AllWalletTransactionHistory - This method returns list all wallet transactions
func (s *Server) AllWalletTransactionHistory(w http.ResponseWriter, r *http.Request) {
	walletTransaction := models.WalletTransaction{}

	//Fetch all wallet transactions
	walletTransactions, err := walletTransaction.AllWalletTransactions(s.DB)
	if err != nil {
		utilities.ERROR(w, http.StatusInternalServerError, err, "")
		return
	}
	utilities.JSON(w, http.StatusOK, walletTransactions, "List of all wallet transactions")
}

//UserWalletTransactionHistory - This method returns list all wallet transactions belonging to a user
func (s *Server) UserWalletTransactionHistory(w http.ResponseWriter, r *http.Request) {
	wallet := models.Wallet{}
	walletTransaction := models.WalletTransaction{}

	//Get Logged in UserID
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		utilities.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"), "Invalid Token")
		return
	}

	//Get user wallet
	userWallet, err := wallet.FindUserWalletByID(userID, s.DB)
	if err != nil {
		utilities.ERROR(w, http.StatusInternalServerError, err, "")
		return
	}

	//Fetch all wallet transactions that belongs to a single user wallet
	walletTransactions, err := walletTransaction.UserWalletTransactions(userWallet.ID, s.DB)
	if err != nil {
		utilities.ERROR(w, http.StatusInternalServerError, err, "")
		return
	}

	utilities.JSON(w, http.StatusOK, walletTransactions, "List of all user wallet transactions")
}

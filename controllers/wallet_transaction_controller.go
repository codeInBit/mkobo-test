package controllers

import (
	"net/http"

	"github.com/codeInBit/mkobo-test/models"
	"github.com/codeInBit/mkobo-test/utilities"
)

//AllWalletTransactionHistory - This method returns list all wallet transactions
func (s *Server) AllWalletTransactionHistory(w http.ResponseWriter, r *http.Request) {
	walletTransaction := models.WalletTransaction{}

	walletTransactions, err := walletTransaction.AllWalletTransfers(s.DB)
	if err != nil {
		utilities.ERROR(w, http.StatusInternalServerError, err, "")
		return
	}
	utilities.JSON(w, http.StatusOK, walletTransactions, "List of all wallet transactions")
}

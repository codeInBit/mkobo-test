package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//Wallet - Wallet struct that represents the Wallet model
type Wallet struct {
	gorm.Model
	UserID            uint `gorm:"not null" json:"user_id"`
	Balance           int  `gorm:"default:10000" json:"balance"`
	WalletTransaction []WalletTransaction
}

//FindUserWalletByID - Returns a user's walllet
func (w *Wallet) FindUserWalletByID(uid uint, db *gorm.DB) (*Wallet, error) {
	var err error
	wallet := Wallet{}

	err = db.Debug().Where("user_id = ?", uid).Take(&wallet).Error
	if err != nil {
		return &Wallet{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Wallet{}, errors.New("Wallet Not Found")
	}
	return &wallet, err
}

//Transfer - Transfer Value between two users
func (w *Wallet) Transfer(senderID, recipientID uint, amount int, db *gorm.DB) (*WalletTransaction, error) {
	var err error
	walletTransaction := WalletTransaction{}

	//Sender wallet
	senderWallet, err := w.FindUserWalletByID(senderID, db)
	if err != nil {
		return nil, err
	}

	//Recipient Wallet
	recipientWallet, err := w.FindUserWalletByID(recipientID, db)
	if err != nil {
		return nil, err
	}

	//Check if sender has enough value to transfer, if not, log the transaction as failed transaction
	if senderWallet.Balance/100 >= amount {
		//Debit recipient
		senderTransaction, err := w.DebitWallet(senderWallet.ID, amount, db)
		if err != nil {
			return nil, err
		}

		//Credit Sender
		w.CreditWallet(recipientWallet.ID, amount, db)

		return senderTransaction, nil
	}

	//Log failed transaction based on insufficient balance
	_, err = walletTransaction.SaveTransaction(senderWallet.ID, amount*100, senderWallet.Balance, senderWallet.Balance, "non", "failure", "Insufficient Balance", db)
	err = errors.New("Transfer cancelled insufficient balance")
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//CreditWallet - This method credits a user's wallet
func (w *Wallet) CreditWallet(walletID uint, amount int, db *gorm.DB) (*WalletTransaction, error) {
	var err error
	walletTransaction := WalletTransaction{}

	//Get wallet
	wallet, err := w.FindUserWalletByID(walletID, db)
	if err != nil {
		return nil, err
	}
	newAmount := amount * 100
	prevBalance := wallet.Balance
	currentBalance := wallet.Balance + newAmount

	// Update wallet balance
	db = db.Debug().Model(&Wallet{}).Where("id = ?", walletID).Take(&Wallet{}).UpdateColumns(
		map[string]interface{}{
			"balance": currentBalance,
		},
	)
	if db.Error != nil {
		return nil, db.Error
	}

	//Log transaction
	transaction, err := walletTransaction.SaveTransaction(walletID, newAmount, prevBalance, currentBalance, "cr", "success", "Wallet credited", db)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

//DebitWallet - This method debits a user's wallet
func (w *Wallet) DebitWallet(walletID uint, amount int, db *gorm.DB) (*WalletTransaction, error) {
	var err error
	walletTransaction := WalletTransaction{}

	//Get wallet
	wallet, err := w.FindUserWalletByID(walletID, db)
	if err != nil {
		return nil, err
	}
	newAmount := amount * 100
	prevBalance := wallet.Balance
	currentBalance := wallet.Balance - newAmount

	// Update wallet balance
	db = db.Debug().Model(&Wallet{}).Where("id = ?", walletID).Take(&Wallet{}).UpdateColumns(
		map[string]interface{}{
			"balance": currentBalance,
		},
	)
	if db.Error != nil {
		return nil, db.Error
	}

	//Log transaction
	transaction, err := walletTransaction.SaveTransaction(walletID, newAmount, prevBalance, currentBalance, "dr", "success", "Wallet debited", db)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

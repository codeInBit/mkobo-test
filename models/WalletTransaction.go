package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

//WalletTransaction - WalletTransaction struct that represents the WalletTransaction model, it logs every transaction that occured in a wallet
type WalletTransaction struct {
	gorm.Model
	WalletID       uint   `gorm:"not null" json:"wallet_id"`
	Amount         int    `gorm:"default:10000" json:"amount"`
	PrevBalance    int    `gorm:"" json:"prev_balance"`
	CurrentBalance int    `gorm:"" json:"current_balance"`
	Effect         string `gorm:"" json:"effect"`
	Status         string `gorm:"" json:"status"`
	Reference      string `gorm:"" json:"reference"`
	Narration      string `gorm:"" json:"narration"`
}

//BeforeSave - This function performs some operation before gorm Create operation
func (wt *WalletTransaction) BeforeSave() error {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	wt.Reference = strconv.Itoa(generator.Int())

	return nil
}

//SaveTransaction - This logs every transaction that happened on every users wallet
func (wt *WalletTransaction) SaveTransaction(walletID uint, amount, prevBalance, currentBalance int, effect, status, narration string, db *gorm.DB) error {
	var err error

	err = db.Debug().Create(&wt).Error
	if err != nil {
		return err
	}

	return nil
}

//AllWalletTransactions - This method fetches all wallet transfer
func (wt *WalletTransaction) AllWalletTransactions(db *gorm.DB) (*[]WalletTransaction, error) {
	var err error
	walletTransactions := []WalletTransaction{}

	err = db.Debug().Model(&WalletTransaction{}).Limit(100).Find(&walletTransactions).Error
	if err != nil {
		return &[]WalletTransaction{}, err
	}
	return &walletTransactions, err
}

//UserWalletTransactions - This method fetches all wallet transfer
func (wt *WalletTransaction) UserWalletTransactions(walletID uint, db *gorm.DB) (*[]WalletTransaction, error) {
	var err error
	walletTransactions := []WalletTransaction{}

	err = db.Debug().Model(&WalletTransaction{}).Where("wallet_id = ?", walletID).Limit(100).Find(&walletTransactions).Error
	if err != nil {
		return &[]WalletTransaction{}, err
	}
	return &walletTransactions, err
}

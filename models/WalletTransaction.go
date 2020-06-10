package models

import (
	"github.com/jinzhu/gorm"
)

//WalletTransaction - WalletTransaction struct that represents the WalletTransaction model, it logs every transaction that occured in a wallet
type WalletTransaction struct {
	gorm.Model
	WalletID       uint   `gorm:"not null" json:"wallet_id"`
	Amount         int    `gorm:"default:10000" json:"amount"`
	PrevBalance    int    `gorm:"" json:"prev_balance"`
	CurrentBalance int    `gorm:"" json:"current_balance"`
	Status         string `gorm:"" json:"status"`
	Reference      string `gorm:"" json:"reference"`
}

package models

import (
	"github.com/jinzhu/gorm"
)

//Wallet - Wallet struct that represents the Wallet model
type Wallet struct {
	gorm.Model
	UserID  uint `gorm:"not null" json:"user_id"`
	Balance int  `gorm:"default:10000" json:"balance"`
}

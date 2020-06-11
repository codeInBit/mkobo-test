package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

//PasswordReset - PasswordReset struct that represents the User model
type PasswordReset struct {
	gorm.Model
	Email string `gorm:"size:100;not null" json:"email"`
	Token string `gorm:"size:255;not null" json:"token"`
}

//BeforeSave - This function performs some operation before gorm Create operation
func (pr *PasswordReset) BeforeSave() error {
	pr.Token = uuid.New().String()
	return nil
}

//SaveResetToken - Save user in database
func (pr *PasswordReset) SaveResetToken(db *gorm.DB) (*PasswordReset, error) {

	var err error
	err = db.Debug().Create(&pr).Error
	if err != nil {
		return &PasswordReset{}, err
	}
	return pr, nil
}

//FindResetRecordByToken - Finds a reset record by tokin, and returns the password reset object
func (pr *PasswordReset) FindResetRecordByToken(token string, db *gorm.DB) (*PasswordReset, error) {
	var err error
	passwordReset := PasswordReset{}

	err = db.Debug().Where("token = ?", token).Take(&passwordReset).Error
	if err != nil {
		return &PasswordReset{}, errors.New("Password reset request record not found")
	}
	if gorm.IsRecordNotFoundError(err) {
		return &PasswordReset{}, errors.New("Password reset request record not found")
	}
	return &passwordReset, err
}

//DeleteAResetRecord - Find a reset record by email and soft delete it
func (pr *PasswordReset) DeleteAResetRecord(email string, db *gorm.DB) (int64, error) {

	db = db.Debug().Where("email = ?", email).Take(&PasswordReset{}).Delete(&PasswordReset{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

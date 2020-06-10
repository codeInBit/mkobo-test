package models

import (
	"errors"
	"html"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/codeInBit/mkobo-test/auth"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User - User struct that represents the User model
type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Wallet   Wallet
}

//Hash - This accepts a password string and returned the hashed version
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword - This compares an hash and a password string
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//BeforeSave - This function performs some operation before gorm Create operation
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

//Prepare - Prepares inputed value
func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

//Validate - performs validation check
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "register":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	case "forgotpassword":
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

//SaveUser - Save user in database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	wallet := Wallet{}

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	//Ceate Wallet
	wallet.UserID = u.ID
	err = db.Debug().Create(&wallet).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

//SignIn - Accept valid email and password to signs a user in, and return an access token
func (u *User) SignIn(email, password string, db *gorm.DB) (string, error) {
	var err error

	user := User{}

	err = db.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}

//FindUserByEmail - Finds a user by email, and returns the user object
func (u *User) FindUserByEmail(email string, db *gorm.DB) (*User, error) {
	var err error
	user := User{}

	err = db.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &user, err
}

//FindUserByID - Finds a user by ID, and returns the user object
func (u *User) FindUserByID(uid uint, db *gorm.DB) (*User, error) {
	var err error
	user := User{}

	err = db.Debug().Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &user, err
}

//ChangePassword - Finds a user by email, and returns the user object
func (u *User) ChangePassword(email, password string, db *gorm.DB) (*User, error) {
	var err error
	user := User{}
	user.Password = password

	// To hash the password
	err = u.BeforeSave()
	if err != nil {
		errors.New("Sorry, An Error occured")
	}

	db = db.Debug().Where("email = ?", email).Take(&user).Update("password", user.Password)

	if db.Error != nil {
		return &User{}, db.Error
	}

	return &user, err
}

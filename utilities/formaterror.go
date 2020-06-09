package utilities

import (
	"errors"
	"strings"
)

//FormatError - Accecpts a string and returns a better formatted string
func FormatError(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}

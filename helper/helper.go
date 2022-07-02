package helper

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// hashing password
func HashingPS(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err.Error())
	}

	return hashedPassword
}

// fungsi pengecekan alamat email
func ValidMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

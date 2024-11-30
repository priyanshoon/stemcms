package server

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password []byte) ([]byte, error) {
	resultPass, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return nil, err
	}
	return resultPass, nil
}

func ComparePass(hashPassword, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashPassword, password); err != nil {
		return false
	}

	return true
}

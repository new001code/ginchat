package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordEncode(p string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("password encoding err: %s", err)
	}
	return string(hash)
}

func PasswordCheck(oldP, newP string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldP), []byte(newP))

	if err != nil {
		return false
	} else {
		return true
	}
}

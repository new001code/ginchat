package util

import (
	"crypto/rand"
	"crypto/rsa"

	"golang.org/x/crypto/bcrypt"
)

var (
	RSAPrivateKey *rsa.PrivateKey
)

func init() {
	RSAPrivateKey = getPrivateKey()
}

func PasswordEncode(p string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		ErrorLogger.Printf("password encoding err: %s", err)
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

func getPrivateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		ErrorLogger.Println("create key pair err:", err)
	}
	return privateKey
}

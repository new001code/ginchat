package util

import (
	"crypto/rand"
	"crypto/rsa"

	"ginchat/common"

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
		common.Logger.Errorf("password encoding err: %s", err)
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
		common.Logger.Errorf("create key pair err: %s", err)
	}
	return privateKey
}

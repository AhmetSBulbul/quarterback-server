package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err.Error()
	}

	return string(hash)
}

func ComparePasswords(hash string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	if err != nil {
		// log.Println(err)
		return false
	}
	return true
}

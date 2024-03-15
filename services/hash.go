package services

import "golang.org/x/crypto/bcrypt"

func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

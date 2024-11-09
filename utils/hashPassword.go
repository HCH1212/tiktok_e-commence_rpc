package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func HashPassword(password string) (string, error) {
	saltedPassword := password
	bytes, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证密码
func VerifyPassword(plainPassword, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}

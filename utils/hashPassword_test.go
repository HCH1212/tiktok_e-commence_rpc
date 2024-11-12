package utils

import (
	"fmt"
	"testing"
)

var (
	secret     = "123456789"
	hashSecret string
	err        error
)

func TestHashPassword(t *testing.T) {
	hashSecret, err = HashPassword(secret)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashSecret)
}

func TestVerifyPassword(t *testing.T) {
	hashSecret, _ = HashPassword(secret)
	res := VerifyPassword(secret, hashSecret)
	fmt.Println(res)
}

package utils

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
)

func ById(id uint) model.User {
	var user model.User
	dao.DB.Where("id = ?", id).First(&user)
	return user
}

func ByEmail(email string) model.User {
	var user model.User
	dao.DB.Where("email = ?", email).First(&user)
	return user
}

package utils

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
)

// user操作
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

// product操作
func BySUK(suk string) model.Product {
	var product model.Product
	dao.DB.Where("sku = ?", suk).First(&product)
	return product
}

func ByProductId(id uint) model.Product {
	var product model.Product
	dao.DB.Where("id = ?", id).First(&product)
	return product
}

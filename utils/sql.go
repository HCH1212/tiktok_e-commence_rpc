package utils

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
)

// user操作
func ById(id uint) (*model.User, error) {
	var user *model.User
	res := dao.DB.Table("users").Where("id=?", id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func ByEmail(email string) (*model.User, error) {
	var user *model.User
	res := dao.DB.Table("users").Where("email=?", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// product操作
func BySUK(suk string) (*model.Product, error) {
	var product *model.Product
	res := dao.DB.Table("products").Where("suk=?", suk).First(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func ByName(name string) ([]*model.Product, error) {
	var products []*model.Product
	res := dao.DB.Table("products").Where("name=?", name).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}

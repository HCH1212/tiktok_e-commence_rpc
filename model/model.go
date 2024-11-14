package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}

type Product struct {
	gorm.Model
	SUK         string   `gorm:"type:varchar(255);unique" json:"suk"`
	Name        string   `gorm:"type:varchar(100)" json:"name"`
	Price       float32  `gorm:"type:int" json:"price"`
	Description string   `gorm:"type:varchar(255)" json:"description"`
	Picture     string   `gorm:"type:varchar(255)" json:"picture"`
	Category    []string `gorm:"type:json" json:"category"`
}

type Cart struct {
	UserId uint64  `gorm:"primaryKey" json:"user_id"`
	ProdID uint    `gorm:"index;" json:"-"` // 添加外键字段ProdID
	Prod   Product `gorm:"foreignKey:ProdID" json:"prod"`
}

type Order struct {
	gorm.Model
	UserId  uint64 `gorm:"primaryKey" json:"user_id"`
	SUK     string `gorm:"type:varchar(255)" json:"suk"`
	Address string `json:"address"`
	IsPay   bool   `gorm:"default:false" json:"is_pay"`
}

type Payment struct {
	gorm.Model
	UserId  uint64  `gorm:"primaryKey" json:"user_id"`
	OrderId uint64  `gorm:"primaryKey" json:"order_id"`
	Amount  float32 `gorm:"type:decimal(10,2)" json:"amount"`
	CardNum string  `gorm:"type:varchar(255)" json:"card_num"`
}

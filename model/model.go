package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
}

package dao

import (
	"fmt"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// 用gorm初始化mysql
func InitMysql() {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	DB = db
	err = DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Cart{}, &model.Order{}) //自动生成表
	if err != nil {
		log.Println(err)
	}
}

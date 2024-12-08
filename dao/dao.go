package dao

import (
	"context"
	"fmt"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

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
		logrus.Println(err)
	}
	if db == nil {
		logrus.Println("db is nil")
	}

	DB = db
	// 链路追踪
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Cart{}, &model.Order{}, &model.Payment{}) //自动生成表
	if err != nil {
		logrus.Println(err)
	}
}

func InitRedis() {
	// 配置 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),     // Redis 地址和端口
		Password: viper.GetString("redis.password"), // 如果没有设置密码，保持为空
		DB:       0,                                 // 使用默认数据库
	})

	// 测试连接
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		logrus.Println("无法连接到 Redis:", err)
		return
	}
	logrus.Println("成功连接到 Redis")

	RDB = rdb
}

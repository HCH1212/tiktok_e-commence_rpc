package email

import (
	"encoding/json"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func InitProducer(msg *model.Email) {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatal(err)
	}
	defer nc.Close()
	data, _ := json.Marshal(msg)
	// 发布消息到主题 "email"
	err = nc.Publish("email", data)
	if err != nil {
		logrus.Fatal(err)
	}
}

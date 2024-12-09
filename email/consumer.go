package email

import (
	"encoding/json"
	"fmt"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func InitConsumer() {
	// 连接到 NATS 服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatal(err)
	}
	defer nc.Close()

	logrus.Info("Connected to NATS")
	// 订阅主题 "email"
	_, err = nc.Subscribe("email", func(msg *nats.Msg) {
		var email model.Email
		_ = json.Unmarshal(msg.Data, &email)
		fmt.Println(email)
	})
	if err != nil {
		logrus.Fatal(err)
	}

	// 保持连接，直到关闭
	select {}
}

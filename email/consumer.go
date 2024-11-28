package email

import (
	"context"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/email"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func consumerInit() {
	sub, err := Nc.Subscribe("email", func(msg *nats.Msg) {
		var req email.EmailReq
		err = proto.Unmarshal(msg.Data, &req)
		if err != nil {
			klog.Error(err)
			return
		}
		var emailImpl EmailImpl
		_, err = emailImpl.Send(context.Background(), &req)
		if err != nil {
			klog.Error(err)
			return
		}
	})
	if err != nil {
		panic(err)
	}
	server.RegisterShutdownHook(func() {
		_ = sub.Unsubscribe()
		Nc.Close()
	})
}

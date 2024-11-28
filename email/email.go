package email

import (
	"context"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/email"
	"github.com/kr/pretty"
)

// 服务段消费消息，客户端生产消息

type EmailImpl struct{}

func (i *EmailImpl) Send(ctx context.Context, req *email.EmailReq) (resp *email.EmailResp, err error) {
	_, err = pretty.Println(req)
	if err != nil {
		return nil, err
	}
	return &email.EmailResp{}, nil
}

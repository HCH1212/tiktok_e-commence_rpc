package payment

import (
	"context"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/payment"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
)

type PaymentImpl struct{}

func (p PaymentImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	pay := &model.Payment{
		UserId:  req.UserId,
		OrderId: req.OrderId,
		Amount:  req.Amount,
		CardNum: req.CardNum,
	}
	dao.DB.Create(pay)
	resp = &payment.ChargeResp{PaymentId: uint64(pay.ID)}
	return
	// 之后记得把订单状态改为已支付
}

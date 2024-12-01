package order

import (
	"context"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/HCH1212/tiktok_e-commence_rpc/utils"
)

// 后面所有订单相关都改为用redis存储，实现自动过期

type OrderImpl struct{}

func (i *OrderImpl) CreateOrder(ctx context.Context, req *order.Order) (resp *order.OrderId, err error) {
	or := &model.Order{
		UserId:  req.UserId,
		SUK:     req.Suk,
		Address: req.Address,
	}
	dao.DB.Create(or)
	resp = &order.OrderId{OrderId: uint64(or.ID)}
	return
}

func (i *OrderImpl) ListOrder(ctx context.Context, req *order.UserId) (resp *order.ListOrderResp, err error) {
	res, err := utils.ByUserIdForOrder(req.UserId)
	if err != nil {
		return nil, err
	}
	ors := make([]*order.Order, len(res))
	for in, v := range res {
		ors[in] = &order.Order{
			UserId:  v.UserId,
			Suk:     v.SUK,
			Address: v.Address,
		}
	}
	resp = &order.ListOrderResp{Orders: ors}
	return
}

func (i *OrderImpl) IsPaidOrder(ctx context.Context, req *order.OrderId) (resp *order.Empty, err error) {
	dao.DB.Table("orders").Where("id=?", req.OrderId).Update("is_pay", true)
	return
}

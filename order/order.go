package order

import (
	"context"
	"encoding/json"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"strconv"
	"time"
)

type OrderImpl struct{}

func (i *OrderImpl) CreateOrder(ctx context.Context, req *order.Order) (resp *order.OrderId, err error) {
	or := &model.Order{
		UserId:  req.UserId,
		SUK:     req.Suk,
		Address: req.Address,
	}
	dao.DB.Create(or)

	// 将订单数据缓存到 Redis
	orderJSON, _ := json.Marshal(or)
	dao.RDB.Set(ctx, "order"+strconv.Itoa(int(or.ID)), orderJSON, time.Hour*24)
	// 将订单 ID 添加到用户的订单列表缓存中
	dao.RDB.RPush(ctx, "user"+strconv.Itoa(int(or.UserId)), strconv.Itoa(int(or.ID)), time.Hour*24)

	resp = &order.OrderId{OrderId: uint64(or.ID)}
	return
}

func (i *OrderImpl) ListOrder(ctx context.Context, req *order.UserId) (resp *order.ListOrderResp, err error) {
	// 从 Redis 获取用户的订单 ID 列表
	orderIDs, err := dao.RDB.LRange(ctx, "user"+strconv.Itoa(int(req.UserId)), 0, -1).Result()
	if err != nil || len(orderIDs) == 0 {
		// 如果缓存中不存在，代表都已过期
		return nil, err
	}

	// 从缓存中获取订单详细信息
	ors := make([]*order.Order, len(orderIDs))
	for j, id := range orderIDs {
		orderJSON, err := dao.RDB.Get(ctx, "order"+id).Result()
		if err != nil {
			continue
		}
		var ord model.Order
		if json.Unmarshal([]byte(orderJSON), &ord) == nil {
			ors[j] = &order.Order{
				UserId:  ord.UserId,
				Suk:     ord.SUK,
				Address: ord.Address,
				IsPay:   ord.IsPay,
			}
		}
	}
	resp = &order.ListOrderResp{Orders: ors}
	return
}

func (i *OrderImpl) IsPaidOrder(ctx context.Context, req *order.OrderId) (resp *order.Empty, err error) {
	dao.DB.Table("orders").Where("id=?", req.OrderId).Update("is_pay", true)
	// 更新缓存
	orderJSON, err := dao.RDB.Get(ctx, "order"+strconv.Itoa(int(req.OrderId))).Result()
	if err == nil {
		var ord model.Order
		if json.Unmarshal([]byte(orderJSON), &ord) == nil {
			ord.IsPay = true
			newOrderJSON, _ := json.Marshal(ord)
			dao.RDB.Set(ctx, "order"+strconv.Itoa(int(req.OrderId)), newOrderJSON, time.Hour*24)
		}
	}
	return
}

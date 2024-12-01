package cart

import (
	"context"
	"encoding/json"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/HCH1212/tiktok_e-commence_rpc/utils"
	"strconv"
	"time"
)

type CartImpl struct{}

func (i *CartImpl) AddItem(ctx context.Context, req *cart.ItemReq) (resp *cart.Empty, err error) {
	// 查询缓存
	var res *model.Product
	cartData, err := dao.RDB.Get(ctx, strconv.Itoa(int(req.UserId))+req.Suk).Result()
	if err == nil && len(cartData) > 0 {
		_ = json.Unmarshal([]byte(cartData), &res)
	} else {
		res, err = utils.BySUK(req.Suk)
		if err != nil {
			return nil, err
		}
	}

	c := &model.Cart{UserId: req.UserId, Prod: *res, ProdID: res.ID}
	dao.DB.Create(c)

	// 更新缓存
	cartJSON, _ := json.Marshal(c)
	dao.RDB.Set(ctx, strconv.Itoa(int(req.UserId))+req.Suk, cartJSON, time.Hour*24)

	return
}

func (i *CartImpl) GetCart(ctx context.Context, req *cart.UserId) (resp *cart.GetCartResp, err error) {
	res, err := utils.ByUserId(req.UserId)
	if err != nil {
		return nil, err
	}
	pros := make([]*cart.Product, len(res))
	for in, v := range res {
		pros[in] = &cart.Product{
			Id:          uint64(v.ProdID),
			SUK:         v.Prod.SUK,
			Name:        v.Prod.Name,
			Price:       v.Prod.Price,
			Description: v.Prod.Description,
			Picture:     v.Prod.Picture,
			Category:    v.Prod.Category,
		}
	}
	resp = &cart.GetCartResp{Products: pros}
	return
}

func (i *CartImpl) DeleteItem(ctx context.Context, req *cart.ItemReq) (resp *cart.Empty, err error) {
	dao.DB.Table("carts").Where("user_id=?", req.UserId).Where("prod.suk=?", req.Suk).Delete(&model.Cart{})

	// 删除缓存
	dao.RDB.Del(ctx, strconv.Itoa(int(req.UserId))+req.Suk)

	return
}

func (i *CartImpl) EmptyCart(ctx context.Context, req *cart.UserId) (resp *cart.Empty, err error) {
	dao.DB.Table("carts").Where("user_id=?", req.UserId).Delete(&model.Cart{})
	return
}

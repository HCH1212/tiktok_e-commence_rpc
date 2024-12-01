package product

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/HCH1212/tiktok_e-commence_rpc/utils"
	"gorm.io/gorm"
	"time"
)

type ProductImpl struct{}

func (i *ProductImpl) CreateProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	// 商品的SUK必需传并且唯一
	// 先看redis中是否有
	res, err := dao.RDB.Get(ctx, req.SUK).Result()
	if err == nil && res != "" {
		return nil, errors.New("商品已存在")
	}

	_, err = utils.BySUK(req.SUK)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pro := model.Product{
			SUK:         req.SUK,
			Name:        req.Name,
			Price:       req.Price,
			Description: req.Description,
			Picture:     req.Picture,
			Category:    req.Category,
		}
		dao.DB.Create(&pro)

		// 添加到缓存
		productJSON, _ := json.Marshal(pro)
		dao.RDB.Set(ctx, pro.SUK, productJSON, time.Hour*24)

		resp = &product.ProductId{Id: uint64(pro.ID)}
		return resp, nil
	} else if err != nil {
		return nil, err
	} else {
		return nil, errors.New("商品已存在")
	}
}

func (i *ProductImpl) ChangeProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	var res *model.Product

	// 先看redis中是否有
	res1, err1 := dao.RDB.Get(ctx, req.SUK).Result()
	if err1 == nil && res1 != "" {
		// 存在
		_ = json.Unmarshal([]byte(res1), &res)
	} else {
		res, err = utils.BySUK(req.SUK)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("商品不存在")
		} else if err != nil {
			return nil, err
		}
	}

	pro := model.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Picture:     req.Picture,
		Category:    req.Category,
	}
	// 更新数据库
	dao.DB.Where("id = ?", res.ID).Save(&pro)

	// 更新缓存
	productJSON, _ := json.Marshal(pro)
	dao.RDB.Set(ctx, pro.SUK, productJSON, time.Hour*24)

	resp = &product.ProductId{Id: uint64(pro.ID)}
	return
}

func (i *ProductImpl) DeleteProduct(ctx context.Context, req *product.ProductId) (resp *product.Pass, err error) {
	var pro model.Product
	// 删除缓存
	// 先通过id查suk，然后删除缓存
	dao.DB.Where("id = ?", req.Id).First(&pro)
	dao.RDB.Del(ctx, pro.SUK)

	dao.DB.Where("id = ?", req.Id).Delete(&pro)

	resp = &product.Pass{Pass: true}
	return
}

func (i *ProductImpl) FindProduct(ctx context.Context, req *product.ProductSUK) (resp *product.Product, err error) {
	// 从缓存中获取
	proData, err := dao.RDB.Get(ctx, req.SUK).Result()
	if err == nil && proData != "" {
		var pro model.Product
		_ = json.Unmarshal([]byte(proData), &pro)

		resp = &product.Product{
			Id:          uint64(pro.ID),
			SUK:         pro.SUK,
			Name:        pro.Name,
			Price:       pro.Price,
			Description: pro.Description,
			Picture:     pro.Picture,
			Category:    pro.Category,
		}
		return resp, nil
	}

	res, err := utils.BySUK(req.SUK)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("商品不存在")
	} else if err != nil {
		return nil, err
	}
	resp = &product.Product{
		Id:          uint64(res.ID),
		SUK:         res.SUK,
		Name:        res.Name,
		Price:       res.Price,
		Description: res.Description,
		Picture:     res.Picture,
		Category:    res.Category,
	}
	return
}

func (i *ProductImpl) FindProducts(ctx context.Context, req *product.SearchReq) (resp *product.SearchResp, err error) {
	// 按名称查询不走缓存
	res, err := utils.ByName(req.Name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("商品不存在")
	} else if err != nil {
		return nil, err
	}
	pros := make([]*product.Product, len(res))
	for in, v := range res {
		pros[in] = &product.Product{
			SUK:         v.SUK,
			Name:        v.Name,
			Price:       v.Price,
			Description: v.Description,
			Picture:     v.Picture,
			Category:    v.Category,
		}
	}
	resp = &product.SearchResp{Products: pros}
	return
}

package product

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product"
	"github.com/HCH1212/tiktok_e-commence_rpc/meili"
	"github.com/HCH1212/tiktok_e-commence_rpc/model"
	"github.com/HCH1212/tiktok_e-commence_rpc/utils"
	"github.com/meilisearch/meilisearch-go"
)

type ProductImpl struct{}

func (i *ProductImpl) CreateProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	// 商品的SUK必需传并且唯一
	res := utils.BySUK(req.SUK)
	if res.ID != 0 {
		return nil, errors.New("商品已存在")
	}
	pro := model.Product{
		SUK:         req.SUK,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Picture:     req.Picture,
		Category:    req.Category,
	}
	dao.DB.Create(&pro)
	resp = &product.ProductId{Id: uint64(pro.ID)}
	return
}

func (i *ProductImpl) ChangeProduct(ctx context.Context, req *product.Product) (resp *product.ProductId, err error) {
	res := utils.BySUK(req.SUK)
	if res.ID == 0 {
		return nil, errors.New("商品不存在")
	}
	pro := model.Product{
		SUK:         req.SUK,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Picture:     req.Picture,
		Category:    req.Category,
	}
	// 更新数据库
	dao.DB.Where("id = ?", res.ID).Save(&pro)
	resp = &product.ProductId{Id: uint64(pro.ID)}
	return
}

func (i *ProductImpl) DeleteProduct(ctx context.Context, req *product.ProductId) (resp *product.Pass, err error) {
	var pro model.Product
	pro.ID = uint(req.Id)
	dao.DB.Where("id = ?", req.Id).Delete(&pro)
	resp = &product.Pass{Pass: true}
	return
}

func (i *ProductImpl) FindProduct(ctx context.Context, req *product.ProductSUK) (resp *product.Product, err error) {
	res := utils.BySUK(req.SUK)
	if res.ID == 0 {
		return nil, errors.New("商品不存在")
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

func (i *ProductImpl) FindProducts(ctx context.Context, req *product.MeiliReq) (resp *product.MeiliResp, err error) {
	// 从数据库取出数据加入到meilisearch
	var products []model.Product
	res := dao.DB.Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	documents := make([]map[string]interface{}, len(products))
	for i, pro := range products {
		documents[i] = map[string]interface{}{
			"suk":         pro.SUK,
			"id":          pro.ID,
			"name":        pro.Name,
			"price":       pro.Price,
			"description": pro.Description,
			"category":    pro.Category,
			"picture":     pro.Picture,
		}
	}
	_, err = meili.Client.Index("id").AddDocuments(documents)
	if err != nil {
		return nil, err
	}

	// 将搜索结构返回为结构体
	searchResults, err := meili.Client.Index("id").Search(req.Query, &meilisearch.SearchRequest{
		AttributesToHighlight: []string{"*"},
	})
	if err != nil {
		return nil, err
	}

	var pros []*product.Product
	for _, hit := range searchResults.Hits {
		prod := &product.Product{
			Id:          hit.(map[string]interface{})["id"].(uint64),
			SUK:         hit.(map[string]interface{})["suk"].(string),
			Name:        hit.(map[string]interface{})["name"].(string),
			Price:       hit.(map[string]interface{})["price"].(int64),
			Picture:     hit.(map[string]interface{})["picture"].(string),
			Description: hit.(map[string]interface{})["description"].(string),
			Category:    hit.(map[string]interface{})["category"].([]string),
		}
		pros = append(pros, prod)
	}

	resp = &product.MeiliResp{Products: pros}
	return
}

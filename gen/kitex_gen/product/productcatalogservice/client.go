// Code generated by Kitex v0.11.3. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	product "github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateProduct(ctx context.Context, Req *product.Product, callOptions ...callopt.Option) (r *product.ProductId, err error)
	ChangeProduct(ctx context.Context, Req *product.Product, callOptions ...callopt.Option) (r *product.ProductId, err error)
	DeleteProduct(ctx context.Context, Req *product.ProductId, callOptions ...callopt.Option) (r *product.Pass, err error)
	FindProduct(ctx context.Context, Req *product.ProductSUK, callOptions ...callopt.Option) (r *product.Product, err error)
	FindProducts(ctx context.Context, Req *product.SearchReq, callOptions ...callopt.Option) (r *product.SearchResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductCatalogServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductCatalogServiceClient struct {
	*kClient
}

func (p *kProductCatalogServiceClient) CreateProduct(ctx context.Context, Req *product.Product, callOptions ...callopt.Option) (r *product.ProductId, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) ChangeProduct(ctx context.Context, Req *product.Product, callOptions ...callopt.Option) (r *product.ProductId, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) DeleteProduct(ctx context.Context, Req *product.ProductId, callOptions ...callopt.Option) (r *product.Pass, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) FindProduct(ctx context.Context, Req *product.ProductSUK, callOptions ...callopt.Option) (r *product.Product, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindProduct(ctx, Req)
}

func (p *kProductCatalogServiceClient) FindProducts(ctx context.Context, Req *product.SearchReq, callOptions ...callopt.Option) (r *product.SearchResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FindProducts(ctx, Req)
}

// Code generated by Kitex v0.11.3. DO NOT EDIT.

package cartservice

import (
	"context"
	cart "github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddItem(ctx context.Context, Req *cart.ItemReq, callOptions ...callopt.Option) (r *cart.Empty, err error)
	GetCart(ctx context.Context, Req *cart.UserId, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	DeleteItem(ctx context.Context, Req *cart.ItemReq, callOptions ...callopt.Option) (r *cart.Empty, err error)
	EmptyCart(ctx context.Context, Req *cart.UserId, callOptions ...callopt.Option) (r *cart.Empty, err error)
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
	return &kCartServiceClient{
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

type kCartServiceClient struct {
	*kClient
}

func (p *kCartServiceClient) AddItem(ctx context.Context, Req *cart.ItemReq, callOptions ...callopt.Option) (r *cart.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddItem(ctx, Req)
}

func (p *kCartServiceClient) GetCart(ctx context.Context, Req *cart.UserId, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCart(ctx, Req)
}

func (p *kCartServiceClient) DeleteItem(ctx context.Context, Req *cart.ItemReq, callOptions ...callopt.Option) (r *cart.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteItem(ctx, Req)
}

func (p *kCartServiceClient) EmptyCart(ctx context.Context, Req *cart.UserId, callOptions ...callopt.Option) (r *cart.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EmptyCart(ctx, Req)
}

// Code generated by Kitex v0.11.3. DO NOT EDIT.

package authservice

import (
	"context"
	auth "github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetToken(ctx context.Context, Req *auth.UserId, callOptions ...callopt.Option) (r *auth.TwoToken, err error)
	ParseAccessToken(ctx context.Context, Req *auth.AccessToken, callOptions ...callopt.Option) (r *auth.UserId, err error)
	ParseRefreshToken(ctx context.Context, Req *auth.RefreshToken, callOptions ...callopt.Option) (r *auth.UserId, err error)
	ExecRefreshToken(ctx context.Context, Req *auth.RefreshToken, callOptions ...callopt.Option) (r *auth.TwoToken, err error)
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
	return &kAuthServiceClient{
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

type kAuthServiceClient struct {
	*kClient
}

func (p *kAuthServiceClient) GetToken(ctx context.Context, Req *auth.UserId, callOptions ...callopt.Option) (r *auth.TwoToken, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetToken(ctx, Req)
}

func (p *kAuthServiceClient) ParseAccessToken(ctx context.Context, Req *auth.AccessToken, callOptions ...callopt.Option) (r *auth.UserId, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ParseAccessToken(ctx, Req)
}

func (p *kAuthServiceClient) ParseRefreshToken(ctx context.Context, Req *auth.RefreshToken, callOptions ...callopt.Option) (r *auth.UserId, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ParseRefreshToken(ctx, Req)
}

func (p *kAuthServiceClient) ExecRefreshToken(ctx context.Context, Req *auth.RefreshToken, callOptions ...callopt.Option) (r *auth.TwoToken, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ExecRefreshToken(ctx, Req)
}

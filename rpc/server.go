package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/user"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func authRpc() server.Server {
	authServer := authservice.NewServer(new(auth.AuthImpl), server.WithRegistry(common()), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "auth",
	}), server.WithReusePort(true))

	return authServer
}

func userRpc() server.Server {
	userServer := userservice.NewServer(new(user.UserImpl), server.WithRegistry(common()), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "user",
	}), server.WithReusePort(true))

	return userServer
}

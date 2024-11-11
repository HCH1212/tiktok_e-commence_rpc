package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/user"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/spf13/viper"
	"log"
)

func InitRpcServer() {
	r, err := consul.NewConsulRegister(viper.GetString("consul.addr"))
	if err != nil {
		log.Fatal(err)
	}

	authServer := authservice.NewServer(new(auth.AuthImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "auth",
	}))
	err = authServer.Run()
	if err != nil {
		log.Fatal(err)
	}

	userServer := userservice.NewServer(new(user.UserImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "user",
	}))
	err = userServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}

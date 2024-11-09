package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/user"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-consul"
	"log"
)

func Auth() {
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	authServer := authservice.NewServer(new(auth.AuthImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "authServer",
	}))
	err = authServer.Run()
	if err != nil {
		log.Fatal(err)
	}

}

func User() {
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	userServer := userservice.NewServer(new(user.UserImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "userServer",
	}))
	err = userServer.Run()
	if err != nil {
		log.Fatal(err)
	}
}

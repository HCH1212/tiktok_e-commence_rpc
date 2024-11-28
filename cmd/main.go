package main

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/config"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/email"
	"github.com/HCH1212/tiktok_e-commence_rpc/rpc"
)

func main() {
	config.InitConfig()
	dao.InitMysql()
	email.Init()
	rpc.InitRpcServer(7)
}

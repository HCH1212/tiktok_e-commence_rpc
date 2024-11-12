package main

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/config"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/meili"
	"github.com/HCH1212/tiktok_e-commence_rpc/rpc"
)

func main() {
	config.InitConfig()
	dao.InitMysql()
	meili.InitMeili()

	rpc.InitRpcServer(3)
}

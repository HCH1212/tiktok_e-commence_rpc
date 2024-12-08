package main

import (
	"context"
	"github.com/HCH1212/tiktok_e-commence_rpc/common/mtl"
	"github.com/HCH1212/tiktok_e-commence_rpc/config"
	"github.com/HCH1212/tiktok_e-commence_rpc/dao"
	"github.com/HCH1212/tiktok_e-commence_rpc/log"
	"github.com/HCH1212/tiktok_e-commence_rpc/rpc"
)

func main() {
	config.InitConfig()
	// 自定义log
	log.InitDefaultLogger()

	// 全局链路追踪
	p := mtl.InitTracing("tiktok_e-commence_rpc")
	defer p.Shutdown(context.Background())

	dao.InitMysql()
	dao.InitRedis()
	rpc.InitRpcServer(6)
}

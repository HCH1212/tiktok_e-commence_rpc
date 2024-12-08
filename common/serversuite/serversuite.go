package serversuite

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 可包装options和middleware

type CommonServerSuite struct {
	CurrentServiceName string
	MetricsPort        string
}

func (s CommonServerSuite) Options() []server.Option {
	register := mtl.InitMetric(s.CurrentServiceName, s.MetricsPort, viper.GetString("consul.addr"))

	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(register))),
		server.WithSuite(tracing.NewServerSuite()),
	}

	r, err := consul.NewConsulRegister(viper.GetString("consul.addr"))
	if err != nil {
		logrus.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	return opts
}

package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/cart"
	"github.com/HCH1212/tiktok_e-commence_rpc/email"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart/cartservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/email/emailservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order/orderservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/payment/paymentservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product/productcatalogservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/order"
	"github.com/HCH1212/tiktok_e-commence_rpc/payment"
	"github.com/HCH1212/tiktok_e-commence_rpc/product"
	"github.com/HCH1212/tiktok_e-commence_rpc/user"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"net"
)

func authRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":8081") // TODO address里面不用加ip，consul有专门的ip
	if err != nil {
		panic(err)
	}
	return authservice.NewServer(new(auth.AuthImpl),
		server.WithRegistry(common()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "auth",
		}),
	)
}

func userRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":8082")
	if err != nil {
		panic(err)
	}
	return userservice.NewServer(new(user.UserImpl),
		server.WithRegistry(common()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "user",
		}),
	) // TODO 用端口复用会出问题 server.WithReusePort(true)

	//2024/11/12 18:38:36.457105 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:51852, error=biz error: 用户已存在
	//2024/11/12 18:38:37.725347 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45904, error=unknown method Register
	//2024/11/12 18:38:38.927851 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45912, error=biz error: 用户已存在
	//2024/11/12 18:38:39.865708 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45926, error=unknown method Register
	//2024/11/12 18:38:40.827435 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45932, error=unknown method Register

}

func productRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":8083")
	if err != nil {
		panic(err)
	}
	return productcatalogservice.NewServer(new(product.ProductImpl),
		server.WithRegistry(common()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "product",
		}),
	)
}

func cartRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":8084")
	if err != nil {
		panic(err)
	}
	return cartservice.NewServer(new(cart.CartImpl),
		server.WithRegistry(common()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "cart",
		}),
	)
}

func orderRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":8085")
	if err != nil {
		panic(err)
	}
	return orderservice.NewServer(new(order.OrderImpl),
		server.WithRegistry(common()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "order",
		}),
	)
}

func paymentRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":8086")
	if err != nil {
		panic(err)
	}
	return paymentservice.NewServer(new(payment.PaymentImpl),
		server.WithRegistry(common()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "payment",
		}),
	)
}

func emailRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":8087")
	if err != nil {
		panic(err)
	}
	return emailservice.NewServer(new(email.EmailImpl),
		server.WithRegistry(common()),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "email",
		}),
	)
}

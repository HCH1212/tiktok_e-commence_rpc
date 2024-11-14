package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/cart"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart/cartservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order/orderservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/payment/paymentservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product/productcatalogservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/order"
	"github.com/HCH1212/tiktok_e-commence_rpc/payment"
	"github.com/HCH1212/tiktok_e-commence_rpc/product"
	"github.com/HCH1212/tiktok_e-commence_rpc/user"
	"github.com/cloudwego/kitex/server"
	"net"
)

//func authRpc() server.Server {
//	authServer := authservice.NewServer(new(auth.AuthImpl), server.WithRegistry(common()), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
//		ServiceName: "auth",
//	}), server.WithReusePort(true)) // TODO 服务注册监听多个端口会出现问题，所以直接复用端口可以解决
//
//	return authServer
//}
//
//func userRpc() server.Server {
//	userServer := userservice.NewServer(new(user.UserImpl), server.WithRegistry(common()), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
//		ServiceName: "user",
//	}), server.WithReusePort(true))
//
//	return userServer
//}

// 暂时不使用consul
func authRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	return authservice.NewServer(new(auth.AuthImpl), server.WithServiceAddr(addr))
}

func userRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	if err != nil {
		panic(err)
	}
	return userservice.NewServer(new(user.UserImpl), server.WithServiceAddr(addr)) // TODO 用端口复用会出问题

	//2024/11/12 18:38:36.457105 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:51852, error=biz error: 用户已存在
	//2024/11/12 18:38:37.725347 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45904, error=unknown method Register
	//2024/11/12 18:38:38.927851 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45912, error=biz error: 用户已存在
	//2024/11/12 18:38:39.865708 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45926, error=unknown method Register
	//2024/11/12 18:38:40.827435 default_server_handler.go:259: [Error] KITEX: processing request error, remoteService=, remoteAddr=127.0.0.1:45932, error=unknown method Register

}

func productRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8083")
	if err != nil {
		panic(err)
	}
	return productcatalogservice.NewServer(new(product.ProductImpl), server.WithServiceAddr(addr))
}

func cartRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8084")
	if err != nil {
		panic(err)
	}
	return cartservice.NewServer(new(cart.CartImpl), server.WithServiceAddr(addr))
}

func orderRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8085")
	if err != nil {
		panic(err)
	}
	return orderservice.NewServer(new(order.OrderImpl), server.WithServiceAddr(addr))
}

func paymentRpc() server.Server {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8086")
	if err != nil {
		panic(err)
	}
	return paymentservice.NewServer(new(payment.PaymentImpl), server.WithServiceAddr(addr))
}

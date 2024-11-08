.PHONY: kitex-auth

kitex-auth:
	@cp idl/auth.proto auth && cd auth && kitex -module github.com/HCH1212/tiktok_e-commence_rpc auth.proto && rm -rf auth.proto

.PHONY: kitex-cart

kitex-cart:
	@cp idl/cart.proto cart && cd cart && kitex -module github.com/HCH1212/tiktok_e-commence_rpc cart.proto && rm -rf cart.proto

.PHONY: kitex-checkout

kitex-checkout:
	@cp idl/checkout.proto checkout && cd checkout && kitex -module github.com/HCH1212/tiktok_e-commence_rpc checkout.proto && rm -rf checkout.proto

.PHONY: kitex-order

kitex-order:
	@cp idl/order.proto order && cd order && kitex -module github.com/HCH1212/tiktok_e-commence_rpc order.proto && rm -rf order.proto

.PHONY: kitex-payment

kitex-payment:
	@cp idl/payment.proto payment && cd payment && kitex -module github.com/HCH1212/tiktok_e-commence_rpc payment.proto && rm -rf payment.proto

.PHONY: kitex-product

kitex-product:
	@cp idl/product.proto product && cd product && kitex -module github.com/HCH1212/tiktok_e-commence_rpc product.proto && rm -rf product.proto

.PHONY: kitex-user

kitex-user:
	@cp idl/user.proto user && cd user && kitex -module github.com/HCH1212/tiktok_e-commence_rpc user.proto && rm -rf user.proto

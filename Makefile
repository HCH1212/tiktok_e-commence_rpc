.PHONY: kitex-auth

kitex-auth:
	@cp idl/auth.proto gen && cd gen && kitex -module github.com/HCH1212/tiktok_e-commence_rpc auth.proto && rm -rf auth.proto

.PHONY: kitex-cart

kitex-cart:
	@cp idl/cart.proto gen && cd gen && kitex -module github.com/HCH1212/tiktok_e-commence_rpc cart.proto && rm -rf cart.proto

.PHONY: kitex-order

kitex-order:
	@cp idl/order.proto gen && cd gen && kitex -module github.com/HCH1212/tiktok_e-commence_rpc order.proto && rm -rf order.proto

.PHONY: kitex-payment

kitex-payment:
	@cp idl/payment.proto gen && cd gen && kitex -module github.com/HCH1212/tiktok_e-commence_rpc payment.proto && rm -rf payment.proto

.PHONY: kitex-product

kitex-product:
	@cp idl/product.proto gen && cd gen && kitex -module github.com/HCH1212/tiktok_e-commence_rpc product.proto && rm -rf product.proto

.PHONY: kitex-user

kitex-user:
	@cp idl/user.proto gen && cd gen && kitex -module github.com/HCH1212/tiktok_e-commence_rpc user.proto && rm -rf user.proto

.PHONY: docker

docker:
	@sudo docker compose up -d

#.PHONY: consul
#
#consul:
#	@open "http://localhost:8500/ui/"

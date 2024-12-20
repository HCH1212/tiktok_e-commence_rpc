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

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: consul
consul:
	@open "http://127.0.0.1:8500/ui/"

.PHONY: dockerfile
dockerfile:
	@go mod vendor && sudo docker build --network host -t your/image:tag .

.PHONY: blackbox
blackbox:
	@cd blackbox_exporter-0.24.0.linux-amd64 && ./blackbox_exporter --config.file="blackbox.yml"

.PHONY: alertmanager
alertmanager:
	@cd alertmanager-0.28.0-rc.0.linux-amd64 && ./alertmanager --config.file="alertmanager.yml"


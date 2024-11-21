FROM golang:1.23.2

WORKDIR /usr/src/tiktok_e-commence_rpc

## 设置代理环境变量
#ENV HTTP_PROXY=http://127.0.0.1:7897
#ENV HTTPS_PROXY=http://127.0.0.1:7897
#ENV NO_PROXY=localhost,127.0.0.1

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

## 本地下载好mod后再复制过来，避免使用代理带来的问题
COPY vendor/ /usr/src/tiktok_e-commence_rpc/vendor/

COPY . .
RUN go build -v -o /opt/tiktok_e-commence_rpc/server ./cmd/main.go

COPY ./config /opt/tiktok_e-commence_rpc/config

WORKDIR /opt/tiktok_e-commence_rpc

EXPOSE 8080

CMD ["./server"]

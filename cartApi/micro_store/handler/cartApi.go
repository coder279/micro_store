package handler

import (
	"context"
    "encoding/json"
	log "github.com/micro/go-micro/v2/logger"
	cart "github.com/coder279/cart"
	cartApi "github.com/coder279/cartApi/proto/cartApi"
)

type CartApi struct{}

// CartApi.Call 通过API向外暴露为/cartApi/call，接收http请求
// 即：/cartApi/call请求会调用go.micro.api.cartApi 服务的CartApi.Call方法
func (e *CartApi) Call(ctx context.Context, req *cartApi.Request, rsp *cartApi.Response) error {
	log.Info("Received CartApi.Call request")
	rsp.StatusCode = 200
    b, _ := json.Marshal("{}")
	rsp.Body = string(b)
	return nil
}

package handler

import (
	"context"
    "encoding/json"
	log "github.com/micro/go-micro/v2/logger"

	paymentApi "github.com/coder279/paymentApi/proto/paymentApi"
)

type PaymentApi struct{}

// PaymentApi.Call 通过API向外暴露为/paymentApi/call，接收http请求
// 即：/paymentApi/call请求会调用go.micro.api.paymentApi 服务的PaymentApi.Call方法
func (e *PaymentApi) Call(ctx context.Context, req *paymentApi.Request, rsp *paymentApi.Response) error {
	log.Info("Received PaymentApi.Call request")
	rsp.StatusCode = 200
    b, _ := json.Marshal("{}")
	rsp.Body = string(b)
	return nil
}

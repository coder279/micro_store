package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/coder279/cart/proto/cart"
	cartApi "github.com/coder279/cartApi/proto/cartApi"
	log "github.com/micro/go-micro/v2/logger"
	"strconv"
)

type CartApi struct{
	CartService go_micro_service_cart.CartService
}

// CartApi.Call 通过API向外暴露为/cartApi/call，接收http请求
// 即：/cartApi/call请求会调用go.micro.api.cartApi 服务的CartApi.Call方法
func (e *CartApi) FindAll(ctx context.Context, req *cartApi.Request, res *cartApi.Response) error {
	log.Info("接受到 /cartApi/findAll 访问请求")
	log.Info(e.CartService)
	if _,ok := req.Get["user_id"]; !ok {
		return errors.New("访问不到user_id")
	}
	userIdString := req.Get["user_id"].Values[0]
	log.Info("到了这部")
	userId,err := strconv.ParseInt(userIdString,10,64)
	if err != nil {
		log.Error(err)
		return err
	}
	//获取购物车所有商品
	cartAll,err := e.CartService.GetAll(context.TODO(),&go_micro_service_cart.CartFindAll{UserId:userId})
	if err != nil {
		log.Error(err)
		return err
	}
    b, err := json.Marshal(cartAll)
    if err != nil {
		log.Error(err)
    	return err
	}
	res.StatusCode = 200
	res.Body = string(b)
	return nil
}

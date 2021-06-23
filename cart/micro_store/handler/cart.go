package handler

import (
	"context"
	"github.com/coder279/cart/common"
	"github.com/coder279/cart/domain/model"
	"github.com/coder279/cart/domain/service"
	"github.com/coder279/cart/proto/cart"
)
type Cart struct{
     CartDataService service.ICartDataService
}
// 添加购物车
func (c *Cart) AddCart(ctx context.Context,req *go_micro_service_cart.CartInfo,res *go_micro_service_cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(req,cart)
	res.CartId,err = c.CartDataService.AddCart(cart)
	return err
}

// 清空购物车
func (c *Cart) CleanCart(ctx context.Context,req *go_micro_service_cart.Clean,res *go_micro_service_cart.Response) (err error){
	if err := c.CartDataService.CleanCart(req.UserId); err != nil {
		return err
	}
	res.Msg = "购物车清空成功"
	return nil
}

// 添加购物车数量
func (c *Cart) Incr(ctx context.Context,req *go_micro_service_cart.Item,res *go_micro_service_cart.Response) (err error){
	if err := c.CartDataService.IncrNum(req.Id,req.ChangeNum); err != nil {
		return err
	}
	res.Msg = "购物车添加成功"
	return nil
}

//减少购物车数量
func (c *Cart) Decr(ctx context.Context,req *go_micro_service_cart.Item,res *go_micro_service_cart.Response) (err error){
	if err := c.CartDataService.DecrNum(req.Id,req.ChangeNum);err != nil {
		return err
	}
	res.Msg = "购物车减少成功"
	return nil
}

func (c *Cart) DeleteItemById(ctx context.Context,req *go_micro_service_cart.CartId,res *go_micro_service_cart.Response)(err error){
	if err := c.CartDataService.DeleteCart(req.Id);err != nil {
		return err
	}
	res.Msg = "购物车删除成功"
	return nil
}
//查询用户所有购物车信息
func (c *Cart) GetAll(ctx context.Context,req *go_micro_service_cart.CartFindAll,res *go_micro_service_cart.CartAll) (err error){
	cartAll,err := c.CartDataService.FindAllCart(req.UserId)
	if err != nil {
		return err
	}
	for _,v := range cartAll{
		cart := &go_micro_service_cart.CartInfo{}
		if err := common.SwapTo(v,cart); err != nil {
			return err
		}
		res.CartInfo = append(res.CartInfo,cart)
	}
	return nil
}
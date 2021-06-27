package handler

import (
	"context"
	"github.com/coder279/order/common"
	"github.com/coder279/order/domain/model"
	"github.com/coder279/order/domain/service"
	"github.com/coder279/order/proto/order"
)
type Order struct{
     OrderDataService service.IOrderDataService
}
//根据订单id查询订单
func (o *Order) GetOrderByID(ctx context.Context,req *go_micro_service_order.OrderID,res *go_micro_service_order.OrderInfo) error {
	order,err := o.OrderDataService.FindOrderByID(req.OrderId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(order,res); err != nil {
		return err
	}
	return nil
}
// 查找所有订单
func (o *Order) GetAllOrder(ctx context.Context,req *go_micro_service_order.AllOrderRequest,res *go_micro_service_order.AllOrder) error {
	orderAll,err := o.OrderDataService.FindAllOrder()
	if err != nil {
		return err
	}
	for _,v := range orderAll{
		order := &go_micro_service_order.OrderInfo{}
		if err := common.SwapTo(v,order); err != nil {
			return err
		}
		res.OrderInfo = append(res.OrderInfo,order)
	}
	return nil
}
//创建订单
func (o *Order) CreateOrder(ctx context.Context,req *go_micro_service_order.OrderInfo,res *go_micro_service_order.OrderID) error {
	orderAdd := &model.Order{}
	if err := common.SwapTo(req,orderAdd);err != nil {
		return err
	}
	orderID,err := o.OrderDataService.AddOrder(orderAdd)
	if err != nil {
		return err
	}
	res.OrderId = orderID
	return nil
}
//删除订单
func (o *Order) DeleteOrderByID(ctx context.Context,req *go_micro_service_order.OrderID,res *go_micro_service_order.Response) error {
	if err := o.OrderDataService.DeleteOrder(req.OrderId);err != nil {
		return err
	}
	res.Msg = "删除成功"
	return nil
}

//更新订单支付状态
func (o *Order) UpdateOrderPayStatus(ctx context.Context,req *go_micro_service_order.PayStatus,res *go_micro_service_order.Response) error {
	if err := o.OrderDataService.UpdatePayStatus(req.OrderId,req.PayStatus); err != nil {
		return err
	}
	res.Msg = "支付成功"
	return nil
}
// 更新物流状态
func (o *Order) UpdateOrderShipStatus(ctx context.Context, req *go_micro_service_order.ShipStatus,res *go_micro_service_order.Response) error {
	if err := o.OrderDataService.UpdateShipStatus(req.OrderId,req.ShipStatus); err != nil {
		return err
	}
	res.Msg = "发货状态成功"
	return nil
}

// 更新订单状态
func (o *Order) UpdateOrder(ctx context.Context,req *go_micro_service_order.OrderInfo,res *go_micro_service_order.Response) error {
	order := &model.Order{}
	if err := common.SwapTo(req,order); err != nil {
		return err
	}
	if err := o.OrderDataService.UpdateOrder(order); err != nil {
		return err
	}
	res.Msg = "订单更新成功"
	return nil
}
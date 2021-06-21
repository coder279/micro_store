package handler

import (
	"context"
	"github.com/coder279/product/common"
	"github.com/coder279/product/domain/model"
	"github.com/coder279/product/domain/service"
	"github.com/coder279/product/proto/product"
)
type Product struct{
     ProductDataService service.IProductDataService
}
//添加商品
func(p *Product) AddProduct(ctx context.Context,req *go_micro_service_product.ProductInfo,res *go_micro_service_product.ResponseProduct)error{
	productAdd := &model.Product{}
	if err := common.SwapTo(req,productAdd);err != nil {
		return err
	}
	productID,err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	res.ProductId = productID
	return nil

}
//根据ID查找商品
func (p *Product) FindProductByID(ctx context.Context,req *go_micro_service_product.RequestID,res *go_micro_service_product.ProductInfo) error {
	productData,err := p.ProductDataService.FindProductByID(req.ProductId)
	if err != nil {
		return err
	}
	err = common.SwapTo(res,productData)
	if err != nil {
		return err
	}
	return nil
}

//修改商品数据
func (p *Product) UpdateProduct(ctx context.Context,req *go_micro_service_product.ProductInfo,res *go_micro_service_product.Response) error{
	productAdd := &model.Product{}
	if err := common.SwapTo(req,productAdd); err != nil {
		return err
	}
	if err := p.ProductDataService.UpdateProduct(productAdd); err != nil {
		return err
	}
	res.Msg = "修改成功"
	return nil
}

//根据商品ID删除
func (p *Product) DeleteProductByID(ctx context.Context,req *go_micro_service_product.RequestID,res *go_micro_service_product.Response) error {
	if err := p.ProductDataService.DeleteProduct(req.ProductId); err != nil {
		return err
	}
	res.Msg = "删除成功"
	return nil
}


//查找所有商品
func (p *Product) FindAllProduct(ctx context.Context,req *go_micro_service_product.RequestAll,res *go_micro_service_product.AllProduct) error {
	productAll,err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	for _,v := range productAll{
		productInfo := &go_micro_service_product.ProductInfo{}
		err := common.SwapTo(v,productInfo)
		if err != nil {
			return err
		}
		res.ProductInfo = append(res.ProductInfo,productInfo)
	}
	return nil
}
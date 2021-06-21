package main

import (
	"context"
	"fmt"
	"github.com/coder279/product/common"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	product "github.com/coder279/product/proto/product"
)

func main(){
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	t,io,err := common.NewTracer("go.micro.service.product.client","localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		//添加注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		)
	productService := product.NewProductService("go.micro.service.product",service.Client())
	productAdd := &product.ProductInfo{
		ProductName: "leo",
		ProductSku:"capS",
		ProductPrice:1.1,
		ProductDescription:"leo-cap",
		ProductCategoryId:1,
		ProductImage: []*product.ProductImage{
			{
				ImageName: "image",
				ImageCode: "0103",
				ImageUrl: "xxx.com",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "size",
				SizeCode: "12",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle: "seo",
			SeoKeywords: "seo",
			SeoDescrption: "seo",
			SeoCode: "seo",
		},

	}
	res,err := productService.AddProduct(context.TODO(),productAdd)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(res)


}
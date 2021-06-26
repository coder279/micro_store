package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/coder279/cart/common"
	go_micro_service_cart "github.com/coder279/cart/proto/cart"
	"github.com/coder279/cartApi/handler"
	cartApi "github.com/coder279/cartApi/proto/cartApi"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context,req client.Request,res interface{},ops ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx,req,res,ops...)
	}, func(e error) error {
		fmt.Println(e)
		return e
	})
}

func NewclientHystrixWrapper() client.Wrapper{
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}

func main() {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.31.112:8500",
		}
	})
	//链路追踪
	t,io,err := common.NewTracer("go.micro.api.cartApi","192.168.31.112:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	//熔断器
	hytrixStramHandler := hystrix.NewStreamHandler()
	hytrixStramHandler.Start()
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("192.168.31.112","9096"),hytrixStramHandler)
		if err != nil {
			log.Error(err)
		}
	}()
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("192.168.31.112:8086"),
		//添加注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(NewclientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
		)

	// Initialise service
	service.Init()
	cartService := go_micro_service_cart.NewCartService("go.micro.service.cart",service.Client())
	log.Info(cartService)
	// Register Handler
	err = cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService:cartService})
	if err != nil {
		log.Error("出现错误")
		log.Error(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

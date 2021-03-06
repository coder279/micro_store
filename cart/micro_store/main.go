package main

import (
	"github.com/coder279/cart/common"
	"github.com/coder279/cart/domain/repository"
	service2 "github.com/coder279/cart/domain/service"
	"github.com/coder279/cart/handler"
	cart "github.com/coder279/cart/proto/cart"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)
var QPS = 100
func main() {
	//配置中心
	consulConfig,err := common.GetConsulConfig("127.0.0.1",8500,"/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.31.112:8500",
		}
	})
	//链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "192.168.31.112:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库连接
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+
		"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SingularTable(true)
	err = repository.NewCartRepository(db).InitTable()
	if err != nil {
		log.Error(err)
	}
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		micro.Address("192.168.31.112:8087"),
		micro.Registry(consul),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialise service
	service.Init()
	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))
	// Register Handler
	cart.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService:cartDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

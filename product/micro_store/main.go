package main

import (
	"github.com/coder279/product/common"
	"github.com/coder279/product/domain/repository"
	service2 "github.com/coder279/product/domain/service"
	"github.com/coder279/product/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"

	product "github.com/coder279/product/proto/product"
)

func main() {
	consulConfig,err :=common.GetConsulConfig("127.0.0.1",8500,"/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	t,io,err := common.NewTracer("go.micro.service.product","localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止创建副表
	db.SingularTable(true)
	//初始化
	//repository.NewProductRepository(db).InitTable()
	productDataService := service2.NewProductDataService(repository.NewProductRepository(db))
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		//添加注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService:productDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

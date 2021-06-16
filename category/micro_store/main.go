package main

import (
	"fmt"
	"github.com/coder279/category/common"
	repository2 "github.com/coder279/category/domain/repository"
	service2 "github.com/coder279/category/domain/service"
	"github.com/coder279/category/handler"
	"github.com/coder279/category/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//配置中心
	consulConfig,err := common.GetConsulConfig("127.0.0.1",8500,"/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)
	//获取mysql配置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	fmt.Println(mysqlInfo)
	//创建数据库连接
	db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)
	rp := repository2.NewCategoryRepository(db)
	rp.InitTable()
	// Initialise service
	service.Init()
	categoryDataService := service2.NewCategoryDataService(repository2.NewCategoryRepository(db))
	err = go_micro_service_category.RegisterCategoryHandler(service.Server(),&handler.Category{CategoryDataService:categoryDataService})
	if err != nil {
		log.Error(err)
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

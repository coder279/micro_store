package main

import (
	"fmt"
	"github.com/coder279/micro_store/domain/repository"
	service2 "github.com/coder279/micro_store/domain/service"
	"github.com/coder279/micro_store/handler"
	"github.com/coder279/micro_store/proto/user"
	"github.com/micro/go-micro/v2"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//service init
	srv.Init()
	//创建数据库
	db,err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:13306)/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)
	rp := repository.NewUserRepository(db)
	rp.InitTable()
	//创建服务实例
	userDataService := service2.NewUserDataService(*repository.NewUserRepository(db))
	//注册handler
	err = go_micro_service_user.RegisterUserHandler(srv.Server(),&handler.User{UserDataService:userDataService})
	if err != nil {
		fmt.Println(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/coder279/micro_store/product/handler"

	product "github.com/coder279/micro_store/product/proto/product"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), new(handler.Product))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

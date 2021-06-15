package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/coder279/micro_store/handler"

	micro_store "github.com/coder279/micro_store/proto/micro_store"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.micro_store"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	micro_store.RegisterMicro_storeHandler(service.Server(), new(handler.Micro_store))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"micro_store/handler"
	pb "micro_store/proto"

	"github.com/micro/micro/v2/service"
	"github.com/micro/micro/v2/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("micro_store"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMicro_storeHandler(srv.Server(), new(handler.Micro_store))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

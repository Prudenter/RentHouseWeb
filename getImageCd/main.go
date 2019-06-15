package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"RentHouseWeb/getImageCd/handler"
	example "RentHouseWeb/getImageCd/proto/example"
	"github.com/micro/go-grpc"
)

func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.getImageCd"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

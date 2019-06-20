package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"RentHouseWeb/postAvatar/handler"
	"RentHouseWeb/postAvatar/subscriber"

	example "RentHouseWeb/postAvatar/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.postAvatar"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.postAvatar", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.postAvatar", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

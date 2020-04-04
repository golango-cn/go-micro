package main

import (
	"go-micro/rpc2http/rpc/handler/greeter"
	"log"
	"github.com/micro/go-micro/v2"
	proto "go-micro/rpc2http/proto"
	"go-micro/rpc2http/rpc/handler/learning"
)

func main() {

	service := micro.NewService(micro.Name("go.micro.api.learning"))

	err := proto.RegisterLearningHandler(service.Server(), new(learning.Handler))
	if err != nil {
		log.Fatal(err)
	}

	err = proto.RegisterGreeterHandler(service.Server(), new(greeter.Handler))
	if err != nil {
		log.Fatal(err)
	}

	service.Init()

	log.Fatal(service.Run())

}

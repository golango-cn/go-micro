package main

import (
	"go-micro/rpc2http/rpc-v2/handler/greeter"
	"log"
	"github.com/micro/go-micro/v2"
	proto "go-micro/rpc2http/proto"
	"go-micro/rpc2http/rpc-v2/handler/learning"
)

func main() {

	service := micro.NewService(micro.Name("go.micro.api.v2.learning"))

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

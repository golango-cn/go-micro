package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"go-micro/micro-stream/client/handler"

	"log"
)

func main() {

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		web.Name("go.micro.golango.web"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8091"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {

			}),
	); err != nil {
		log.Fatal(err)
	}

	service.HandleFunc("/stream1", handler.Stream1)
	service.HandleFunc("/stream2", handler.Stream2)
	service.HandleFunc("/stream3", handler.Stream3)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 2379)}
}

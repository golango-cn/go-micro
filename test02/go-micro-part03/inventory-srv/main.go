package main

import (
	"fmt"
	"go-micro/go-micro-part03/basic/config"
	"go-micro/go-micro-part03/inventory-srv/handler"
	"go-micro/go-micro-part03/basic"
	"go-micro/go-micro-part03/inventory-srv/model"
	"go-micro/go-micro-part03/proto/inventory"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.inventory"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			model.Init()
			return nil
		}),
	)

	// 注册服务
	inventory.RegisterInventoryHandler(service.Server(), new(handler.InventoryHandler))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


func registryOptions(ops *registry.Options) {
	host := config.AppConfig.App.Etcd.Host
	port := config.AppConfig.App.Etcd.Port
	ops.Addrs = []string{fmt.Sprintf("%s:%d", host, port)}
}

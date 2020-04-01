package main

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro/go-micro-part02/auth/model"
	"go-micro/go-micro-part02/basic"
	"go-micro/go-micro-part02/auth/handler"
	"github.com/micro/cli/v2"
	"go-micro/go-micro-part02/basic/config"
	"go-micro/go-micro-part02/basic/db"
	auth  "go-micro/go-micro-part02/proto"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	// 初始化redis
	db.InitRedis()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化handler
			model.Init()
			return nil
		}),
	)

	// 注册服务
	auth.RegisterTokenServiceHandler(service.Server(), new(handler.AuthHandler))

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
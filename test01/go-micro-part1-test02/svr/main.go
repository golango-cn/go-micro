package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/util/log"
	"go-micro/go-micro-part1-test02/svr/basic"
	"go-micro/go-micro-part1-test02/svr/basic/config"
	"go-micro/go-micro-part1-test02/svr/handler"

	s "go-micro/go-micro-part1-test02/svr/proto"
)

func main(){

	basic.Init()


	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.department"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			return nil
		}),
	)

	// 注册服务
	s.RegisterDepartmentServiceHandler(service.Server(), new(handler.DepartmentHander))

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


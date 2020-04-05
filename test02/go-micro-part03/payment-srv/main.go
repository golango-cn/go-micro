package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro/go-micro-part03/basic"
	"go-micro/go-micro-part03/basic/config"
	"go-micro/go-micro-part03/payment-srv/handler"
	"go-micro/go-micro-part03/payment-srv/model"

	payment "go-micro/go-micro-part03/proto/payment"
)

func main() {

	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.payment"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()

			return nil
		}),
	)

	// 注册服务
	payment.RegisterPaymentHandler(service.Server(), new(handler.PaymentHander))

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

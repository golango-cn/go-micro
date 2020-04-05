package main

import (
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro/go-micro-part03/basic"
	"go-micro/go-micro-part03/basic/config"
	"go-micro/go-micro-part03/order-srv/handler"
	"go-micro/go-micro-part03/order-srv/subscriber"
	order "go-micro/go-micro-part03/proto/order"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.orders"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {

			// 初始化handler
			handler.Init()
			// 初始化sub
			subscriber.Init()

			return nil
		}),
	)

	// 侦听订单支付消息
	//err := micro.RegisterSubscriber("mu.micro.book.topic.payment.done", service.Server(), subscriber.PayOrder)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 注册服务
	err := order.RegisterOrdersHandler(service.Server(), new(handler.OrderHandler))
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}


func registryOptions(ops *registry.Options) {
	host := config.AppConfig.App.Etcd.Host
	port := config.AppConfig.App.Etcd.Port
	ops.Addrs = []string{fmt.Sprintf("%s:%d", host, port)}
}
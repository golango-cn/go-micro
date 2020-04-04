package main

import (
	"context"
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"go-micro/go-micro-part02/auth/handler"
	"go-micro/go-micro-part02/auth/model"
	"go-micro/go-micro-part02/basic"
	"go-micro/go-micro-part02/basic/config"
	"go-micro/go-micro-part02/basic/db"
	auth "go-micro/go-micro-part02/proto"

	//"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	// 初始化redis
	db.InitRedis()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建broker
	bk1 := rabbitmq.NewBroker(
		broker.Addrs("amqp://test01:test01@140.143.0.152:5672"),
	)
	//bk1 := broker.NewBroker(
	//	broker.Addrs(fmt.Sprintf("%s:%d", "127.0.0.1", 11089)),
	//)
	bk1.Init()
	bk1.Connect()

	//订阅主题1
	_, err := bk1.Subscribe("mu.micro.book.topic.auth.tokenExpired", func(event broker.Event) error {
		fmt.Println(string(event.Message().Body))
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
		micro.Broker(
			bk1,
		),
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

	// register subscriber with queue, each message is delivered to a unique subscriber
	//micro.RegisterSubscriber("mu.micro.book.topic.auth.tokenExpired", service.Server(), subEv)


	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


// Alternatively a function can be used
func subEv(ctx context.Context,msg interface{}) error {
	log.Infof("[pubsub.2] Received event %#v\n", msg)
	// do something with event
	return nil
}



func registryOptions(ops *registry.Options) {
	host := config.AppConfig.App.Etcd.Host
	port := config.AppConfig.App.Etcd.Port
	ops.Addrs = []string{fmt.Sprintf("%s:%d", host, port)}
}
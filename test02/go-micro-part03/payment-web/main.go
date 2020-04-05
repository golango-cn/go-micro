package main

import (
        "fmt"
        "github.com/micro/cli/v2"
        log "github.com/micro/go-micro/v2/logger"
        "github.com/micro/go-micro/v2/registry"
        "github.com/micro/go-micro/v2/registry/etcd"
        "go-micro/go-micro-part03/basic"
        "go-micro/go-micro-part03/basic/config"
        "net/http"
        "github.com/micro/go-micro/v2/web"
        "payment-web/handler"
)

func main() {
        // 初始化配置
        basic.Init()

        // 使用etcd注册
        micReg := etcd.NewRegistry(registryOptions)

        // 创建新服务
        service := web.NewService(
                web.Name("mu.micro.book.web.payment"),
                web.Version("latest"),
                web.Registry(micReg),
                web.Address(":8090"),
        )

        // 初始化服务
        if err := service.Init(
                web.Action(
                        func(c *cli.Context) {
                                // 初始化handler
                                handler.Init()
                        }),
        ); err != nil {
                log.Fatal(err)
        }

        // 新建订单接口
        authHandler := http.HandlerFunc(handler.PayOrder)
        service.Handle("/payment/pay-order", handler.AuthWrapper(authHandler))

        // 运行服务
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}


func registryOptions(ops *registry.Options) {
        host := config.AppConfig.App.Etcd.Host
        port := config.AppConfig.App.Etcd.Port
        ops.Addrs = []string{fmt.Sprintf("%s:%d", host, port)}
}
package main

import (
        "fmt"

        "go-micro/go-micro-part1-test02/web/basic"

        "github.com/micro/cli/v2"
        log "github.com/micro/go-micro/v2/logger"
        "github.com/micro/go-micro/v2/registry"
        "github.com/micro/go-micro/v2/registry/etcd"
        "github.com/micro/go-micro/v2/web"
        "go-micro/go-micro-part1-test02/web/basic/config"
        "go-micro/go-micro-part1-test02/web/handler"
)

func main() {
        // 初始化配置
        basic.Init()

        // 使用etcd注册
        micReg := etcd.NewRegistry(registryOptions)

        // 创建新服务
        service := web.NewService(
                // 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
                web.Name("mu.micro.book.web.department"),
                web.Version("latest"),
                web.Registry(micReg),
                web.Address(":8087"),
        )

        // 初始化服务
        if err := service.Init(
                web.Action(func(c *cli.Context) {
                        // 初始化handler
                        handler.Init()
                }),
        ); err != nil {
                log.Fatal(err)
        }

        // 注册登录接口
        service.HandleFunc("/department/departments", handler.GetDepartments)

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

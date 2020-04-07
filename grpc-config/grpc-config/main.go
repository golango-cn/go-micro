package main

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"go-micro/grpc-config/grpc-config/basic"
	"go-micro/grpc-config/grpc-config/basic/common"
	"go-micro/grpc-config/grpc-config/basic/config"
	"time"

	_ "github.com/micro/go-micro/v2/config"
	_ "github.com/micro/go-micro/v2/logger"
)

var appName = "auth_srv"
var cfg = &authCfg{}

func main() {
	initCfg()

	fmt.Println("cfg", cfg)

	ch := make(chan int)
	<-ch
}

type authCfg struct {
	common.AppCfg
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("grpc1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		log.Info(err)
	}

	go func() {
		ticker := time.NewTicker(time.Second * 3)

		for {
			<-ticker.C
			log.Infof("[initCfg] 配置，cfg：%v", cfg)
		}

	}()
	log.Infof("[initCfg] 配置，cfg：%v", cfg)

	return
}

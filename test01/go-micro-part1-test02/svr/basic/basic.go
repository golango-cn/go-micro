package basic

import (
	"go-micro/go-micro-part1-test02/svr/basic/config"
	"go-micro/go-micro-part1-test02/svr/basic/db"
)

func Init(){
	config.Init()
	db.Init()
}

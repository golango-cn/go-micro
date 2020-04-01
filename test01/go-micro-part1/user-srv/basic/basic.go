package basic

import (
	"go-micro-part1/user-srv/basic/config"
	"go-micro-part1/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}

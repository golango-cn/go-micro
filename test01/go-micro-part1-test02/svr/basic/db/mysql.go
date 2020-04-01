package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part1-test02/svr/basic/config"
)



func initMysql() {
	var err error

	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.AppConfig.App.Mysql.Url)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.AppConfig.App.Mysql.MaxOpenConnection)

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.AppConfig.App.Mysql.MaxIdleConnection)

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

}

package db

import (
	"database/sql"
	"github.com/go-redis/redis"
)

var (
	mysqlDB *sql.DB
)

func GetDb() *sql.DB {
	return mysqlDB
}

func GetRedis() *redis.Client {
	return client
}

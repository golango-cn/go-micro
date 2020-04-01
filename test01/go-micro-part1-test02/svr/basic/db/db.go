package db

import "database/sql"

var (
	mysqlDB *sql.DB
)

func Init()  {
	initMysql()
}

func GetDb() *sql.DB{
	return mysqlDB
}
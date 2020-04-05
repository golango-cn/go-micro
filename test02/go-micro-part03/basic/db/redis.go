package db

import (
	"github.com/go-redis/redis"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/basic/config"
	"sync"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

// Init 初始化Redis
func InitRedis() {

	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("已经初始化过Redis...")
		return
	}

	initSingle()

	log.Info("初始化Redis，检测连接...")

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("初始化Redis，检测连接Ping.")
	log.Info("初始化Redis，检测连接Ping..")
	log.Infof("初始化Redis，检测连接Ping... %s", pong)

}

func initSingle() {

	conn := config.AppConfig.Redis.Conn
	password := config.AppConfig.Redis.Password
	db := config.AppConfig.Redis.DbNum

	client = redis.NewClient(&redis.Options{
		Addr:     conn,
		Password: password, // no password set
		DB:       db,    // use default DB
	})
}

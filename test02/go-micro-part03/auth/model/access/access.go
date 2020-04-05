package access

import (
	"github.com/go-redis/redis"
	"go-micro/go-micro-part03/basic/db"
	"time"
)

type TokenService struct{

}

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	// tokenIDKeyPrefix tokenID 前缀
	tokenIDKeyPrefix = "token:auth:id:"

	tokenExpiredTopic = "mu.micro.book.topic.auth.tokenExpired"

	// Redis 客户端
	ca  *redis.Client

	tokenService *TokenService
)


// Subject token 持有者
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

func Init()  {
	ca = db.GetRedis()
	tokenService = &TokenService{}
}

func GetTokenService() *TokenService {
	return tokenService
}
package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
	"go-micro/go-micro-part03/basic/config"
	"strconv"
)

// MakeAccessToken 生成token并保存到redis
func (s *TokenService) MakeAccessToken(subject *Subject) (ret string, err error) {

	_, err = s.createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token Claim 失败，err: %s", err)
	}

	// 创建
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)

	token := jwt.New(jwt.GetSigningMethod(jwt.SigningMethodHS256.Name))
	claims := token.Claims.(jwt.MapClaims)


	claims["UserId"], _ = strconv.ParseUint(subject.ID, 10, 64)
	claims["UserName"] = subject.Name

	ret, err = token.SignedString([]byte(config.AppConfig.Jwt.SecretKey))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token失败，err: %s", err)
	}

	// 保存到redis
	err = s.saveTokenToCache(subject, ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 保存token到缓存失败，err: %s", err)
	}

	return
}

// GetCachedAccessToken 获取token
func (s *TokenService) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] 从缓存获取token失败，err: %s", err)
	}

	return
}

// DelUserAccessToken 清除用户token
func (s *TokenService) DelUserAccessToken(tk string) (err error) {
	// 解析token字符串
	claims, err := s.parseToken(tk)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
	}

	// 通过解析到的用户id删除
	err = s.delTokenFromCache(&Subject{
		ID: claims.Subject,
	})

	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 清除用户token，err: %s", err)
	}

	fmt.Println("广播消息", claims.Subject)
	// 广播删除
	msg := &broker.Message{
		Body: []byte(claims.Subject),
	}
	//if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
	//	log.Infof("[pub] 发布token删除消息失败： %v", err)
	//} else {
	//	fmt.Println("[pub] 发布token删除消息：", string(msg.Body))
	//}

	//bk1 := broker.NewBroker(
	//	broker.Addrs(fmt.Sprintf("%s:%d", "127.0.0.1", 11089)),
	//)
	bk1 := rabbitmq.NewBroker(
		broker.Addrs("amqp://test01:test01@140.143.0.152:5672"),
	)
	bk1.Init()
	bk1.Connect()

	if err := bk1.Publish(tokenExpiredTopic, msg); err != nil {
		log.Infof("[pub] 发布token删除消息失败： %v", err)
	} else {
		fmt.Println("[pub] 发布token删除消息：", string(msg.Body))
	}

	return
}

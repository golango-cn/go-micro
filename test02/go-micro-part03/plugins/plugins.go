package plugins

import (
	_ "github.com/dgrijalva/jwt-go"
	_ "go-micro/go-micro-part03/plugins/session"
	_ "go-micro/go-micro-part03/plugins/zap"
	_ "github.com/afex/hystrix-go/hystrix"
)

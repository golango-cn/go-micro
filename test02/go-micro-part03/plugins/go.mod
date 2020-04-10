module go-micro/go-micro-part03/plugins

go 1.14

replace go-micro/go-micro-part03/basic => ../basic

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.2.0
	go-micro/go-micro-part03/basic v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.14.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

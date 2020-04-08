module go-micro/go-micro-part03/auth

go 1.14

replace go-micro/go-micro-part03/proto => ../proto

replace go-micro/go-micro-part03/basic => ../basic

replace go-micro/go-micro-part03/plugins => ../plugins

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.3.0
	go-micro/go-micro-part03/basic v0.0.0-00010101000000-000000000000
	go-micro/go-micro-part03/plugins v0.0.0-00010101000000-000000000000
	go-micro/go-micro-part03/proto v0.0.0-00010101000000-000000000000
	gopkg.in/square/go-jose.v2 v2.3.1
)

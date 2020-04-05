module go-micro/go-micro-part03/web

go 1.14

replace go-micro/go-micro-part03/basic => ../basic

replace go-micro/go-micro-part03/proto => ../proto

replace go-micro/go-micro-part03/plugins => ../plugins

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.4.0
	go-micro/go-micro-part03/basic v0.0.0-00010101000000-000000000000
	go-micro/go-micro-part03/plugins v0.0.0-00010101000000-000000000000
	go-micro/go-micro-part03/proto v0.0.0-00010101000000-000000000000
)

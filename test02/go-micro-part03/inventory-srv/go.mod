module go-micro/go-micro-part03/inventory-srv

go 1.14

replace go-micro/go-micro-part03/basic => ../basic

replace go-micro/go-micro-part03/proto => ../proto

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/protoc-gen-micro v1.0.0
	go-micro/go-micro-part03/basic v0.0.0-00010101000000-000000000000
	go-micro/go-micro-part03/proto v0.0.0-00010101000000-000000000000
)

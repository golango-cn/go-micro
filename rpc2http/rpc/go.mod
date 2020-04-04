module go-micro/rpc2http/rpc

go 1.14

replace go-micro/rpc2http/proto => ../proto

require (
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/protoc-gen-micro v1.0.0
	go-micro/rpc2http/proto v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200222125558-5a598a2470a0
)

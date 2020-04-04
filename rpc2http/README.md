#### 说明
Rpc服务转换成Http服务接口

####

进入proto目录执行
```
protoc --proto_path=. --go_out=. --micro_out=. */*.proto
```

启动API：

```
micro api --handler=rpc --namespace=go.micro.api
```

启动rpc服务
```
go run rpc/main.go
```

启动rpc服务v2
```
go run rpc-v2/main.go
```

执行request.http内的方法
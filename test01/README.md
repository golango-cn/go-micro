
### 源项目
https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro

### SQL表

```sql
CREATE TABLE `_user` (
	`user_id` int ( 11 ) NOT NULL AUTO_INCREMENT,
	`user_name` varchar ( 255 ) NOT NULL,
	`user_pwd` varchar ( 255 ) NOT NULL,
	`created_time` int ( 11 ) DEFAULT NULL,
	`updated_time` int ( 11 ) DEFAULT NULL,
	PRIMARY KEY ( `user_id` ),
UNIQUE KEY `uk_user_name` ( `user_name` ( 191 ) ) USING BTREE 
) ENGINE = InnoDB AUTO_INCREMENT = 10002 DEFAULT CHARSET = utf8mb4;
```

### 启动API
```
micro --registry=etcd --api_namespace=mu.micro.book.web  api --handler=web
```

### 启动user-svr
```
go run main.go plugin.go 
```

### 启动user-web
```
go run main.go 
```


### 测试
```
curl --request POST   --url http://127.0.0.1:8080/user/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'userName=micro&pwd=123'

{"data":{"id":10001,"name":"micro"},"ref":1583578133854912943,"success":true}

```

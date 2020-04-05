运行api
```cassandraql
micro --registry=etcd --api_namespace=mu.micro.book.web  api --handler=web
```

运行auth、svr、web

```cassandraql
curl --request POST   --url http://127.0.0.1:8080/employee/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'name=Parto&pwd=1234'
```

```cassandraql
{"baseResponse":{"isSuccess":true,"errorMessage":"测试成功"},"employee":{"empno":10003,"birthDate":"1959-12-03T00:00:00+08:00","firstName":"Parto","lastName":"Bamford","gendeM","hireDate":"1986-08-28T00:00:00+09:00"},"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODg1OTUyODQsImp0aSI6IlBhcnRvIiwiaWF0IjoxNTg2MDAzMjg0LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTg2MDAzMjg0LCJzdWIiOiJQYXJ0byJ9.S4WaHFs3xvpDxQ5G0ZHjX7QQ7TI2hA-atQh0K8L9ea8"}
```


```cassandraql
curl --request POST \
  --url http://127.0.0.1:8080/employee/logout \
  --cookie 'remember-me-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODg1OTUyODQsImp0aSI6IlBhcnRvIiwiaWF0IjoxNTg2MDAzMjg0LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTg2MDAzMjg0LCJzdWIiOiJQYXJ0byJ9.S4WaHFs3xvpDxQ5G0ZHjX7QQ7TI2hA-atQh0K8L9ea8'
```